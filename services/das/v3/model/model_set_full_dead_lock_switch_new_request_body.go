package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetFullDeadLockSwitchNewRequestBody 设置完整死锁开关请求体
type SetFullDeadLockSwitchNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 开关状态
	SwitchOn bool `json:"switch_on"`
}

func (o SetFullDeadLockSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetFullDeadLockSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetFullDeadLockSwitchNewRequestBody", string(data)}, " ")
}
