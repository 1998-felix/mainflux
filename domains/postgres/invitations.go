// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/absmach/supermq/domains"
	repoerr "github.com/absmach/supermq/pkg/errors/repository"
	"github.com/absmach/supermq/pkg/postgres"
)

func (repo domainRepo) SaveInvitation(ctx context.Context, invitation domains.Invitation) (err error) {
	q := `INSERT INTO invitations (invited_by, user_id, domain_id, role_id, created_at)
		VALUES (:invited_by, :user_id, :domain_id, :role_id, :created_at)`

	dbInv := toDBInvitation(invitation)
	if _, err = repo.db.NamedExecContext(ctx, q, dbInv); err != nil {
		return postgres.HandleError(repoerr.ErrCreateEntity, err)
	}

	return nil
}

func (repo domainRepo) RetrieveInvitation(ctx context.Context, userID, domainID string) (domains.Invitation, error) {
	q := `SELECT invited_by, user_id, domain_id, role_id, created_at, updated_at, confirmed_at, rejected_at FROM invitations WHERE user_id = :user_id AND domain_id = :domain_id;`

	dbinv := dbInvitation{
		UserID:   userID,
		DomainID: domainID,
	}
	rows, err := repo.db.NamedQueryContext(ctx, q, dbinv)
	if err != nil {
		return domains.Invitation{}, postgres.HandleError(repoerr.ErrViewEntity, err)
	}
	defer rows.Close()

	dbinv = dbInvitation{}
	if rows.Next() {
		if err = rows.StructScan(&dbinv); err != nil {
			return domains.Invitation{}, postgres.HandleError(repoerr.ErrViewEntity, err)
		}

		return toInvitation(dbinv), nil
	}

	return domains.Invitation{}, repoerr.ErrNotFound
}

func (repo domainRepo) RetrieveAllInvitations(ctx context.Context, pm domains.InvitationPageMeta) (domains.InvitationPage, error) {
	query := pageQuery(pm)

	q := fmt.Sprintf("SELECT invited_by, user_id, domain_id, role_id, created_at, updated_at, confirmed_at, rejected_at FROM invitations %s LIMIT :limit OFFSET :offset;", query)

	rows, err := repo.db.NamedQueryContext(ctx, q, pm)
	if err != nil {
		return domains.InvitationPage{}, postgres.HandleError(repoerr.ErrViewEntity, err)
	}
	defer rows.Close()

	var items []domains.Invitation
	for rows.Next() {
		var dbinv dbInvitation
		if err = rows.StructScan(&dbinv); err != nil {
			return domains.InvitationPage{}, postgres.HandleError(repoerr.ErrViewEntity, err)
		}
		items = append(items, toInvitation(dbinv))
	}

	tq := fmt.Sprintf(`SELECT COUNT(*) FROM invitations %s`, query)

	total, err := postgres.Total(ctx, repo.db, tq, pm)
	if err != nil {
		return domains.InvitationPage{}, postgres.HandleError(repoerr.ErrViewEntity, err)
	}

	invPage := domains.InvitationPage{
		Total:       total,
		Offset:      pm.Offset,
		Limit:       pm.Limit,
		Invitations: items,
	}

	return invPage, nil
}

func (repo domainRepo) UpdateConfirmation(ctx context.Context, invitation domains.Invitation) (err error) {
	q := `UPDATE invitations SET confirmed_at = :confirmed_at, updated_at = :updated_at WHERE user_id = :user_id AND domain_id = :domain_id`

	dbinv := toDBInvitation(invitation)
	result, err := repo.db.NamedExecContext(ctx, q, dbinv)
	if err != nil {
		return postgres.HandleError(repoerr.ErrUpdateEntity, err)
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return repoerr.ErrNotFound
	}

	return nil
}

func (repo domainRepo) UpdateRejection(ctx context.Context, invitation domains.Invitation) (err error) {
	q := `UPDATE invitations SET rejected_at = :rejected_at, updated_at = :updated_at WHERE user_id = :user_id AND domain_id = :domain_id`

	dbInv := toDBInvitation(invitation)
	result, err := repo.db.NamedExecContext(ctx, q, dbInv)
	if err != nil {
		return postgres.HandleError(repoerr.ErrUpdateEntity, err)
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return repoerr.ErrNotFound
	}

	return nil
}

func (repo domainRepo) DeleteInvitation(ctx context.Context, userID, domain string) (err error) {
	q := `DELETE FROM invitations WHERE user_id = $1 AND domain_id = $2`

	result, err := repo.db.ExecContext(ctx, q, userID, domain)
	if err != nil {
		return postgres.HandleError(repoerr.ErrRemoveEntity, err)
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return repoerr.ErrNotFound
	}

	return nil
}

func pageQuery(pm domains.InvitationPageMeta) string {
	var query []string
	var emq string
	if pm.DomainID != "" {
		query = append(query, "domain_id = :domain_id")
	}
	if pm.UserID != "" {
		query = append(query, "user_id = :user_id")
	}
	if pm.InvitedBy != "" {
		query = append(query, "invited_by = :invited_by")
	}
	if pm.RoleID != "" {
		query = append(query, "role_id = :role_id")
	}
	if pm.InvitedByOrUserID != "" {
		query = append(query, "(invited_by = :invited_by_or_user_id OR user_id = :invited_by_or_user_id)")
	}
	if pm.State == domains.Accepted {
		query = append(query, "confirmed_at IS NOT NULL")
	}
	if pm.State == domains.Pending {
		query = append(query, "confirmed_at IS NULL AND rejected_at IS NULL")
	}
	if pm.State == domains.Rejected {
		query = append(query, "rejected_at IS NOT NULL")
	}

	if len(query) > 0 {
		emq = fmt.Sprintf("WHERE %s", strings.Join(query, " AND "))
	}

	return emq
}

type dbInvitation struct {
	InvitedBy   string       `db:"invited_by"`
	UserID      string       `db:"user_id"`
	DomainID    string       `db:"domain_id"`
	RoleID      string       `db:"role_id,omitempty"`
	Relation    string       `db:"relation"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at,omitempty"`
	ConfirmedAt sql.NullTime `db:"confirmed_at,omitempty"`
	RejectedAt  sql.NullTime `db:"rejected_at,omitempty"`
}

func toDBInvitation(inv domains.Invitation) dbInvitation {
	var updatedAt, confirmedAt, rejectedAt sql.NullTime
	if inv.UpdatedAt != (time.Time{}) {
		updatedAt = sql.NullTime{Time: inv.UpdatedAt, Valid: true}
	}
	if inv.ConfirmedAt != (time.Time{}) {
		confirmedAt = sql.NullTime{Time: inv.ConfirmedAt, Valid: true}
	}
	if inv.RejectedAt != (time.Time{}) {
		rejectedAt = sql.NullTime{Time: inv.RejectedAt, Valid: true}
	}

	return dbInvitation{
		InvitedBy:   inv.InvitedBy,
		UserID:      inv.UserID,
		DomainID:    inv.DomainID,
		RoleID:      inv.RoleID,
		CreatedAt:   inv.CreatedAt,
		UpdatedAt:   updatedAt,
		ConfirmedAt: confirmedAt,
		RejectedAt:  rejectedAt,
	}
}

func toInvitation(dbinv dbInvitation) domains.Invitation {
	var updatedAt, confirmedAt, rejectedAt time.Time
	if dbinv.UpdatedAt.Valid {
		updatedAt = dbinv.UpdatedAt.Time
	}
	if dbinv.ConfirmedAt.Valid {
		confirmedAt = dbinv.ConfirmedAt.Time
	}
	if dbinv.RejectedAt.Valid {
		rejectedAt = dbinv.RejectedAt.Time
	}

	return domains.Invitation{
		InvitedBy:   dbinv.InvitedBy,
		UserID:      dbinv.UserID,
		DomainID:    dbinv.DomainID,
		RoleID:      dbinv.RoleID,
		CreatedAt:   dbinv.CreatedAt,
		UpdatedAt:   updatedAt,
		ConfirmedAt: confirmedAt,
		RejectedAt:  rejectedAt,
	}
}
