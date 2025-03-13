package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityApprovalsResponse Response Object
type ListSecurityApprovalsResponse struct {

	// 工单列表
	Approvals *[]PermissionApprovalOpenapiDto `json:"approvals,omitempty"`

	// 规则组总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecurityApprovalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityApprovalsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityApprovalsResponse", string(data)}, " ")
}
