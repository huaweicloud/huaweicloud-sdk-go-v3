package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskCheckSettingsItem struct {
	// 检查参数对应的key值

	CfgKey string `json:"cfg_key"`
	// 参数状态

	Status string `json:"status"`
	// 检查参数值

	CfgValue *string `json:"cfg_value,omitempty"`
}

func (o TaskCheckSettingsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskCheckSettingsItem struct{}"
	}

	return strings.Join([]string{"TaskCheckSettingsItem", string(data)}, " ")
}
