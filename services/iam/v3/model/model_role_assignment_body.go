package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RoleAssignmentBody struct {
	User *RoleUserAssignmentId `json:"user,omitempty" xml:"user"`

	Role *RoleAssignmentId `json:"role,omitempty" xml:"role"`

	Group *RoleGroupAssignmentId `json:"group,omitempty" xml:"group"`

	Agency *RoleAgencyAssignmentId `json:"agency,omitempty" xml:"agency"`

	Scope *RoleAssignmentScope `json:"scope,omitempty" xml:"scope"`

	// 是否基于所有项目授权。
	IsInherited *bool `json:"is_inherited,omitempty" xml:"is_inherited"`
}

func (o RoleAssignmentBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleAssignmentBody struct{}"
	}

	return strings.Join([]string{"RoleAssignmentBody", string(data)}, " ")
}
