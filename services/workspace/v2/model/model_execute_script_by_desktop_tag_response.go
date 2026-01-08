package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScriptByDesktopTagResponse Response Object
type ExecuteScriptByDesktopTagResponse struct {

	// 任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteScriptByDesktopTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptByDesktopTagResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScriptByDesktopTagResponse", string(data)}, " ")
}
