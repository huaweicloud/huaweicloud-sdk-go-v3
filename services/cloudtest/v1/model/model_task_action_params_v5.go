package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskActionParamsV5 struct {

	// 启停、调试动作（1为启动，0为停止，2为调试）
	ActionId *int32 `json:"action_id,omitempty"`

	// 环境Id
	EnvironmentGroupId *string `json:"environment_group_id,omitempty"`

	// 测试计划Id
	PlanId *string `json:"plan_id,omitempty"`

	// 任务id列表信息
	TaskIds *[]string `json:"taskIds,omitempty"`
}

func (o TaskActionParamsV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskActionParamsV5 struct{}"
	}

	return strings.Join([]string{"TaskActionParamsV5", string(data)}, " ")
}
