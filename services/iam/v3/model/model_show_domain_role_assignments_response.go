package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDomainRoleAssignmentsResponse struct {

	// 返回授权记录的总条数。
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num"`

	RoleAssignments *[]RoleAssignmentBody `json:"role_assignments,omitempty" xml:"role_assignments"`
	HttpStatusCode  int                   `json:"-"`
}

func (o ShowDomainRoleAssignmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainRoleAssignmentsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainRoleAssignmentsResponse", string(data)}, " ")
}
