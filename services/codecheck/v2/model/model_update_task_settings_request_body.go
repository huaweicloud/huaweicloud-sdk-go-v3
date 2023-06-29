package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskSettingsRequestBody 任务配置高级选项的请求信息
type UpdateTaskSettingsRequestBody struct {

	// 高级选项参数的相关信息
	TaskAdvancedSettings []TaskAdvancedSettingsItem `json:"task_advanced_settings"`
}

func (o UpdateTaskSettingsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSettingsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTaskSettingsRequestBody", string(data)}, " ")
}
