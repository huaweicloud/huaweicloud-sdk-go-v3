package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaskSettingsRequestBody **参数解释：** 高级配置参数的相关信息，数组不能为空。 **约束限制：** 不涉及。
type UpdateTaskSettingsRequestBody struct {

	// **参数解释：** 高级配置参数的相关信息，数组不能为空。 **约束限制：** 不涉及。
	TaskAdvancedSettings []TaskAdvancedSettingsItem `json:"task_advanced_settings"`
}

func (o UpdateTaskSettingsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskSettingsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTaskSettingsRequestBody", string(data)}, " ")
}
