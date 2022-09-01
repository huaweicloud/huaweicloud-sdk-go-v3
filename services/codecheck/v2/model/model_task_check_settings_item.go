package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskCheckSettingsItem struct {

	// 检查参数对应的key值
	CfgKey string `json:"cfg_key" xml:"cfg_key"`

	// 参数状态
	Status string `json:"status" xml:"status"`

	// 检查参数值
	CfgValue *string `json:"cfg_value,omitempty" xml:"cfg_value"`
}

func (o TaskCheckSettingsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskCheckSettingsItem struct{}"
	}

	return strings.Join([]string{"TaskCheckSettingsItem", string(data)}, " ")
}
