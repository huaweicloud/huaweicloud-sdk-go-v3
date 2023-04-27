package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadSpecialConfigurationSettingRequest struct {

	// 任务id
	TaskId string `json:"task_id"`

	Body *ConfigurationRequestBody `json:"body,omitempty"`
}

func (o UploadSpecialConfigurationSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSpecialConfigurationSettingRequest struct{}"
	}

	return strings.Join([]string{"UploadSpecialConfigurationSettingRequest", string(data)}, " ")
}
