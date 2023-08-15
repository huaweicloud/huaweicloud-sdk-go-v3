package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigSettingRequest Request Object
type ShowConfigSettingRequest struct {

	// 任务id
	TaskId string `json:"task_id"`

	// 具体请求配置项
	ConfigKey *string `json:"config_key,omitempty"`
}

func (o ShowConfigSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigSettingRequest", string(data)}, " ")
}
