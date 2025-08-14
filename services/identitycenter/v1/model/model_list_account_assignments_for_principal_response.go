package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAssignmentsForPrincipalResponse Response Object
type ListAccountAssignmentsForPrincipalResponse struct {

	// 满足查询条件的账号分配列表
	AccountAssignments *[]AccountAssignmentDto `json:"account_assignments,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAccountAssignmentsForPrincipalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentsForPrincipalResponse struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentsForPrincipalResponse", string(data)}, " ")
}
