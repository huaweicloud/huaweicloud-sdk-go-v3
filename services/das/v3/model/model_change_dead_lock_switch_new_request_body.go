package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDeadLockSwitchNewRequestBody 修改死锁开关请求体
type ChangeDeadLockSwitchNewRequestBody struct {

	// 数据库引擎类型，取值范围：mysql
	EngineType string `json:"engine_type"`

	// 开关状态，取值范围：false（关闭）、true（开启）
	SwitchOn bool `json:"switch_on"`

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 保存时长
	RetentionHours int32 `json:"retention_hours"`
}

func (o ChangeDeadLockSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDeadLockSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeDeadLockSwitchNewRequestBody", string(data)}, " ")
}
