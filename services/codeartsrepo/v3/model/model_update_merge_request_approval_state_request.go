package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMergeRequestApprovalStateRequest Request Object
type UpdateMergeRequestApprovalStateRequest struct {

	// 仓库的主键id
	RepositoryId string `json:"repository_id"`

	// 合并请求的短id
	MergeRequestIid int32 `json:"merge_request_iid"`

	Body *ApprovalActionTypeDto `json:"body,omitempty"`
}

func (o UpdateMergeRequestApprovalStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestApprovalStateRequest struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestApprovalStateRequest", string(data)}, " ")
}
