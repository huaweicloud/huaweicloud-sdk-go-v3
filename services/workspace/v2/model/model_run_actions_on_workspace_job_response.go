package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunActionsOnWorkspaceJobResponse Response Object
type RunActionsOnWorkspaceJobResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunActionsOnWorkspaceJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunActionsOnWorkspaceJobResponse struct{}"
	}

	return strings.Join([]string{"RunActionsOnWorkspaceJobResponse", string(data)}, " ")
}
