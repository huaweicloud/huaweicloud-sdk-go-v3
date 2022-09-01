package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RoleAssignmentScope struct {
	Project *RoleProjectAssignmentId `json:"project,omitempty" xml:"project"`

	Domain *RoleDomainAssignmentId `json:"domain,omitempty" xml:"domain"`

	EnterpriseProject *RoleEnterpriseProjectAssignmentId `json:"enterprise_project,omitempty" xml:"enterprise_project"`
}

func (o RoleAssignmentScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleAssignmentScope struct{}"
	}

	return strings.Join([]string{"RoleAssignmentScope", string(data)}, " ")
}
