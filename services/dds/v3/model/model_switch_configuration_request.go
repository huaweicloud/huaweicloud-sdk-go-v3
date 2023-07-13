package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchConfigurationRequest Request Object
type SwitchConfigurationRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationRequestBody `json:"body,omitempty"`
}

func (o SwitchConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchConfigurationRequest struct{}"
	}

	return strings.Join([]string{"SwitchConfigurationRequest", string(data)}, " ")
}
