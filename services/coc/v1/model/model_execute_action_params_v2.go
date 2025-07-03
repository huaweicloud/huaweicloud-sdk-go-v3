package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteActionParamsV2 struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 动作名称
	Action string `json:"action"`

	// 动作信息
	Params *interface{} `json:"params,omitempty"`
}

func (o ExecuteActionParamsV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteActionParamsV2 struct{}"
	}

	return strings.Join([]string{"ExecuteActionParamsV2", string(data)}, " ")
}
