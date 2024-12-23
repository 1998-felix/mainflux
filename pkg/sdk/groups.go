// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	apiutil "github.com/absmach/supermq/api/http/util"
	"github.com/absmach/supermq/pkg/errors"
)

const (
	groupsEndpoint   = "groups"
	childrenEndpoint = "children"
	MaxLevel         = uint64(5)
	MinLevel         = uint64(1)
)

// Group represents the group of Clients.
// Indicates a level in tree hierarchy. Root node is level 1.
// Path in a tree consisting of group IDs
// Paths are unique per owner.
type Group struct {
	ID                        string    `json:"id,omitempty"`
	DomainID                  string    `json:"domain_id,omitempty"`
	ParentID                  string    `json:"parent_id,omitempty"`
	Name                      string    `json:"name,omitempty"`
	Description               string    `json:"description,omitempty"`
	Metadata                  Metadata  `json:"metadata,omitempty"`
	Level                     int       `json:"level,omitempty"`
	Path                      string    `json:"path,omitempty"`
	Children                  []*Group  `json:"children,omitempty"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
	UpdatedBy                 string    `json:"updated_by,omitempty"`
	Status                    string    `json:"status,omitempty"`
	RoleID                    string    `json:"role_id,omitempty"`
	RoleName                  string    `json:"role_name,omitempty"`
	Actions                   []string  `json:"actions,omitempty"`
	AccessType                string    `json:"access_type,omitempty"`
	AccessProviderId          string    `json:"access_provider_id,omitempty"`
	AccessProviderRoleId      string    `json:"access_provider_role_id,omitempty"`
	AccessProviderRoleName    string    `json:"access_provider_role_name,omitempty"`
	AccessProviderRoleActions []string  `json:"access_provider_role_actions,omitempty"`
}

func (sdk mgSDK) CreateGroup(g Group, domainID, token string) (Group, errors.SDKError) {
	data, err := json.Marshal(g)
	if err != nil {
		return Group{}, errors.NewSDKError(err)
	}
	url := fmt.Sprintf("%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint)

	_, body, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusCreated)
	if sdkerr != nil {
		return Group{}, sdkerr
	}

	g = Group{}
	if err := json.Unmarshal(body, &g); err != nil {
		return Group{}, errors.NewSDKError(err)
	}

	return g, nil
}

func (sdk mgSDK) Groups(pm PageMetadata, domainID, token string) (GroupsPage, errors.SDKError) {
	endpoint := fmt.Sprintf("%s/%s", domainID, groupsEndpoint)
	url, err := sdk.withQueryParams(sdk.groupsURL, endpoint, pm)
	if err != nil {
		return GroupsPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return GroupsPage{}, sdkerr
	}

	gp := GroupsPage{}
	if err := json.Unmarshal(body, &gp); err != nil {
		return GroupsPage{}, errors.NewSDKError(err)
	}

	return gp, nil
}

func (sdk mgSDK) Group(id, domainID, token string) (Group, errors.SDKError) {
	if id == "" {
		return Group{}, errors.NewSDKError(apiutil.ErrMissingID)
	}

	url := fmt.Sprintf("%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id)

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return Group{}, sdkerr
	}

	var t Group
	if err := json.Unmarshal(body, &t); err != nil {
		return Group{}, errors.NewSDKError(err)
	}

	return t, nil
}

func (sdk mgSDK) UpdateGroup(g Group, domainID, token string) (Group, errors.SDKError) {
	data, err := json.Marshal(g)
	if err != nil {
		return Group{}, errors.NewSDKError(err)
	}

	if g.ID == "" {
		return Group{}, errors.NewSDKError(apiutil.ErrMissingID)
	}
	url := fmt.Sprintf("%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, g.ID)

	_, body, sdkerr := sdk.processRequest(http.MethodPut, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return Group{}, sdkerr
	}

	g = Group{}
	if err := json.Unmarshal(body, &g); err != nil {
		return Group{}, errors.NewSDKError(err)
	}

	return g, nil
}

func (sdk mgSDK) SetGroupParent(id, domainID, groupID, token string) errors.SDKError {
	scpg := groupParentReq{ParentID: groupID}
	data, err := json.Marshal(scpg)
	if err != nil {
		return errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, parentEndpoint)
	_, _, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusOK)

	return sdkerr
}

func (sdk mgSDK) RemoveGroupParent(id, domainID, groupID, token string) errors.SDKError {
	rcpg := groupParentReq{ParentID: groupID}
	data, err := json.Marshal(rcpg)
	if err != nil {
		return errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, parentEndpoint)
	_, _, sdkerr := sdk.processRequest(http.MethodDelete, url, token, data, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) AddChildren(id, domainID string, groupIDs []string, token string) errors.SDKError {
	acg := childrenGroupsReq{ChildrenIDs: groupIDs}
	data, err := json.Marshal(acg)
	if err != nil {
		return errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, childrenEndpoint)
	_, _, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusOK)

	return sdkerr
}

func (sdk mgSDK) RemoveChildren(id, domainID string, groupIDs []string, token string) errors.SDKError {
	rcg := childrenGroupsReq{ChildrenIDs: groupIDs}
	data, err := json.Marshal(rcg)
	if err != nil {
		return errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, childrenEndpoint)
	_, _, sdkerr := sdk.processRequest(http.MethodDelete, url, token, data, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) RemoveAllChildren(id, domainID, token string) errors.SDKError {
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, childrenEndpoint, "all")
	_, _, sdkerr := sdk.processRequest(http.MethodDelete, url, token, nil, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) Children(id, domainID string, pm PageMetadata, token string) (GroupsPage, errors.SDKError) {
	endpoint := fmt.Sprintf("%s/%s/%s/%s", domainID, groupsEndpoint, id, childrenEndpoint)
	url, err := sdk.withQueryParams(sdk.groupsURL, endpoint, pm)
	if err != nil {
		return GroupsPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return GroupsPage{}, sdkerr
	}

	gp := GroupsPage{}
	if err := json.Unmarshal(body, &gp); err != nil {
		return GroupsPage{}, errors.NewSDKError(err)
	}

	return gp, nil
}

func (sdk mgSDK) EnableGroup(id, domainID, token string) (Group, errors.SDKError) {
	return sdk.changeGroupStatus(id, enableEndpoint, domainID, token)
}

func (sdk mgSDK) DisableGroup(id, domainID, token string) (Group, errors.SDKError) {
	return sdk.changeGroupStatus(id, disableEndpoint, domainID, token)
}

func (sdk mgSDK) changeGroupStatus(id, status, domainID, token string) (Group, errors.SDKError) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, status)

	_, body, sdkerr := sdk.processRequest(http.MethodPost, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return Group{}, sdkerr
	}
	g := Group{}
	if err := json.Unmarshal(body, &g); err != nil {
		return Group{}, errors.NewSDKError(err)
	}

	return g, nil
}

func (sdk mgSDK) DeleteGroup(id, domainID, token string) errors.SDKError {
	if id == "" {
		return errors.NewSDKError(apiutil.ErrMissingID)
	}
	url := fmt.Sprintf("%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id)
	_, _, sdkerr := sdk.processRequest(http.MethodDelete, url, token, nil, nil, http.StatusNoContent)
	return sdkerr
}

func (sdk mgSDK) Hierarchy(id, domainID string, pm PageMetadata, token string) (GroupsHierarchyPage, errors.SDKError) {
	endpoint := fmt.Sprintf("%s/%s/%s/hierarchy", domainID, groupsEndpoint, id)
	url, err := sdk.withQueryParams(sdk.groupsURL, endpoint, pm)
	if err != nil {
		return GroupsHierarchyPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return GroupsHierarchyPage{}, sdkerr
	}

	hp := GroupsHierarchyPage{}
	if err := json.Unmarshal(body, &hp); err != nil {
		return GroupsHierarchyPage{}, errors.NewSDKError(err)
	}

	return hp, nil
}

func (sdk mgSDK) CreateGroupRole(id, domainID string, rq RoleReq, token string) (Role, errors.SDKError) {
	data, err := json.Marshal(rq)
	if err != nil {
		return Role{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint)
	_, body, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusCreated)
	if sdkerr != nil {
		return Role{}, sdkerr
	}

	role := Role{}
	if err := json.Unmarshal(body, &role); err != nil {
		return Role{}, errors.NewSDKError(err)
	}

	return role, nil
}

func (sdk mgSDK) GroupRoles(id, domainID string, pm PageMetadata, token string) (RolesPage, errors.SDKError) {
	endpoint := fmt.Sprintf("%s/%s/%s/%s", domainID, groupsEndpoint, id, rolesEndpoint)
	url, err := sdk.withQueryParams(sdk.groupsURL, endpoint, pm)
	if err != nil {
		return RolesPage{}, errors.NewSDKError(err)
	}

	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return RolesPage{}, sdkerr
	}

	var rp RolesPage
	if err := json.Unmarshal(body, &rp); err != nil {
		return RolesPage{}, errors.NewSDKError(err)
	}

	return rp, nil
}

func (sdk mgSDK) GroupRole(id, roleName, domainID, token string) (Role, errors.SDKError) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName)
	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return Role{}, sdkerr
	}

	var role Role
	if err := json.Unmarshal(body, &role); err != nil {
		return Role{}, errors.NewSDKError(err)
	}

	return role, nil
}

func (sdk mgSDK) UpdateGroupRole(id, roleName, newName, domainID string, token string) (Role, errors.SDKError) {
	ucr := updateRoleNameReq{Name: newName}
	data, err := json.Marshal(ucr)
	if err != nil {
		return Role{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName)
	_, body, sdkerr := sdk.processRequest(http.MethodPut, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return Role{}, sdkerr
	}

	role := Role{}
	if err := json.Unmarshal(body, &role); err != nil {
		return Role{}, errors.NewSDKError(err)
	}

	return role, nil
}

func (sdk mgSDK) DeleteGroupRole(id, roleName, domainID, token string) errors.SDKError {
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName)
	_, _, sdkerr := sdk.processRequest(http.MethodDelete, url, token, nil, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) AddGroupRoleActions(id, roleName, domainID string, actions []string, token string) ([]string, errors.SDKError) {
	acra := roleActionsReq{Actions: actions}
	data, err := json.Marshal(acra)
	if err != nil {
		return []string{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName, actionsEndpoint)
	_, body, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return []string{}, sdkerr
	}

	res := roleActionsRes{}
	if err := json.Unmarshal(body, &res); err != nil {
		return []string{}, errors.NewSDKError(err)
	}

	return res.Actions, nil
}

func (sdk mgSDK) GroupRoleActions(id, roleName, domainID string, token string) ([]string, errors.SDKError) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName, actionsEndpoint)
	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return nil, sdkerr
	}

	res := roleActionsRes{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, errors.NewSDKError(err)
	}

	return res.Actions, nil
}

func (sdk mgSDK) RemoveGroupRoleActions(id, roleName, domainID string, actions []string, token string) errors.SDKError {
	rcra := roleActionsReq{Actions: actions}
	data, err := json.Marshal(rcra)
	if err != nil {
		return errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName, actionsEndpoint, "delete")
	_, _, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) RemoveAllGroupRoleActions(id, roleName, domainID, token string) errors.SDKError {
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName, actionsEndpoint, "delete-all")
	_, _, sdkerr := sdk.processRequest(http.MethodPost, url, token, nil, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) AddGroupRoleMembers(id, roleName, domainID string, members []string, token string) ([]string, errors.SDKError) {
	acrm := roleMembersReq{Members: members}
	data, err := json.Marshal(acrm)
	if err != nil {
		return []string{}, errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName, membersEndpoint)
	_, body, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusOK)
	if sdkerr != nil {
		return []string{}, sdkerr
	}

	res := roleMembersRes{}
	if err := json.Unmarshal(body, &res); err != nil {
		return []string{}, errors.NewSDKError(err)
	}

	return res.Members, nil
}

func (sdk mgSDK) GroupRoleMembers(id, roleName, domainID string, pm PageMetadata, token string) (RoleMembersPage, errors.SDKError) {
	endpoint := fmt.Sprintf("%s/%s/%s/%s/%s/%s", domainID, groupsEndpoint, id, rolesEndpoint, roleName, membersEndpoint)
	url, err := sdk.withQueryParams(sdk.groupsURL, endpoint, pm)
	if err != nil {
		return RoleMembersPage{}, errors.NewSDKError(err)
	}
	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return RoleMembersPage{}, sdkerr
	}

	res := RoleMembersPage{}
	if err := json.Unmarshal(body, &res); err != nil {
		return RoleMembersPage{}, errors.NewSDKError(err)
	}

	return res, nil
}

func (sdk mgSDK) RemoveGroupRoleMembers(id, roleName, domainID string, members []string, token string) errors.SDKError {
	rcrm := roleMembersReq{Members: members}
	data, err := json.Marshal(rcrm)
	if err != nil {
		return errors.NewSDKError(err)
	}

	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName, membersEndpoint, "delete")
	_, _, sdkerr := sdk.processRequest(http.MethodPost, url, token, data, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) RemoveAllGroupRoleMembers(id, roleName, domainID, token string) errors.SDKError {
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, id, rolesEndpoint, roleName, membersEndpoint, "delete-all")
	_, _, sdkerr := sdk.processRequest(http.MethodPost, url, token, nil, nil, http.StatusNoContent)

	return sdkerr
}

func (sdk mgSDK) AvailableGroupRoleActions(domainID, token string) ([]string, errors.SDKError) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s", sdk.groupsURL, domainID, groupsEndpoint, rolesEndpoint, "available-actions")
	_, body, sdkerr := sdk.processRequest(http.MethodGet, url, token, nil, nil, http.StatusOK)
	if sdkerr != nil {
		return nil, sdkerr
	}

	res := availableRoleActionsRes{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, errors.NewSDKError(err)
	}

	return res.AvailableActions, nil
}
