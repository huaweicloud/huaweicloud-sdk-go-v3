package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDesktopPoolScriptResponse Response Object
type ExecuteDesktopPoolScriptResponse struct {

	// 任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteDesktopPoolScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDesktopPoolScriptResponse struct{}"
	}

	return strings.Join([]string{"ExecuteDesktopPoolScriptResponse", string(data)}, " ")
}
