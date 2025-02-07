// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package authn

import (
	"context"
)

type TokenType uint32

const (
	// AccessToken represents token generated by user.
	AccessToken TokenType = iota
	// PersonalAccessToken represents token generated by user for automation.
	PersonalAccessToken
)

func (t TokenType) String() string {
	switch t {
	case AccessToken:
		return "access token"
	case PersonalAccessToken:
		return "pat"
	default:
		return "unknown"
	}
}

type Session struct {
	Type       TokenType
	PatID      string
	UserID     string
	DomainID   string
	SuperAdmin bool
}

// Authn is supermq authentication library.
//
//go:generate mockery --name Authentication --output=./mocks --filename authn.go --quiet --note "Copyright (c) Abstract Machines"
type Authentication interface {
	Authenticate(ctx context.Context, token string) (Session, error)
}
