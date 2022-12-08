package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UnlockWorkspaceResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnlockWorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockWorkspaceResponse struct{}"
	}

	return strings.Join([]string{"UnlockWorkspaceResponse", string(data)}, " ")
}
