package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMergeRequestApprovalStateResponse Response Object
type UpdateMergeRequestApprovalStateResponse struct {
	Result *ApproverBasicDto `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMergeRequestApprovalStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestApprovalStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestApprovalStateResponse", string(data)}, " ")
}
