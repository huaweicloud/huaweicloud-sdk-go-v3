package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryWorkspacesApprovalResponse Response Object
type ListFactoryWorkspacesApprovalResponse struct {

	// 审批详细信息。
	JobApplySearchList *[]ListFactoryWorkspacesApprovalRespJobApplySearchList `json:"job_apply_search_list,omitempty"`

	// 审批总数量。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFactoryWorkspacesApprovalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryWorkspacesApprovalResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryWorkspacesApprovalResponse", string(data)}, " ")
}
