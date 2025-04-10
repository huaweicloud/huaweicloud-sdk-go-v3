package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityTableApproversResponse Response Object
type ListSecurityTableApproversResponse struct {

	// 审批人列表
	Approvers *[]TableApprover `json:"approvers,omitempty"`

	// 是否已经有权限
	HasTableRule   *bool `json:"has_table_rule,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListSecurityTableApproversResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityTableApproversResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityTableApproversResponse", string(data)}, " ")
}
