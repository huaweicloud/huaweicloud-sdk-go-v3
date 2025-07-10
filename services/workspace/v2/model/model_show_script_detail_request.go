package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScriptDetailRequest Request Object
type ShowScriptDetailRequest struct {

	// 脚本ID。
	ScriptId string `json:"script_id"`

	// 执行脚本的任务ID。
	ScriptTaskId *string `json:"script_task_id,omitempty"`
}

func (o ShowScriptDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowScriptDetailRequest", string(data)}, " ")
}
