package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunActionsOnWorkspaceJobRequest Request Object
type RunActionsOnWorkspaceJobRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	Body *JobActionsReq `json:"body,omitempty"`
}

func (o RunActionsOnWorkspaceJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunActionsOnWorkspaceJobRequest struct{}"
	}

	return strings.Join([]string{"RunActionsOnWorkspaceJobRequest", string(data)}, " ")
}
