package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptExecuteParam 脚本执行相关参数
type ScriptExecuteParam struct {

	// 是否资源受限，true表示不受限，false表示受限
	Resourceful *bool `json:"resourceful,omitempty"`

	// 超时时间，单位：秒，取值范围待定，5 < timeout < 1800
	Timeout int32 `json:"timeout"`

	// 成功率，支持小数点后一位
	SuccessRate float64 `json:"success_rate"`

	// 脚本执行用户，如：root
	ExecuteUser string `json:"execute_user"`

	// 脚本入参列表
	ScriptParams *[]ScriptExecuteInputParam `json:"script_params,omitempty"`
}

func (o ScriptExecuteParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptExecuteParam struct{}"
	}

	return strings.Join([]string{"ScriptExecuteParam", string(data)}, " ")
}
