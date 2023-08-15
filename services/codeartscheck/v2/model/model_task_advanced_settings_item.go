package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskAdvancedSettingsItem struct {

	// 高级选项对应的名称
	Key string `json:"key"`

	// 高级选项对应的取值
	Value string `json:"value"`
}

func (o TaskAdvancedSettingsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAdvancedSettingsItem struct{}"
	}

	return strings.Join([]string{"TaskAdvancedSettingsItem", string(data)}, " ")
}
