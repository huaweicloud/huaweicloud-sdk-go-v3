package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchConfigurationRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id" xml:"config_id"`

	Body *ApplyConfigurationRequestBody `json:"body,omitempty" xml:"body"`
}

func (o SwitchConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchConfigurationRequest struct{}"
	}

	return strings.Join([]string{"SwitchConfigurationRequest", string(data)}, " ")
}
