package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConfigurationRequest struct {
	// 参数模板ID。

	ConfigId string `json:"config_id"`

	Body *UpdateConfigurationRequestBody `json:"body,omitempty"`
}

func (o UpdateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationRequest", string(data)}, " ")
}
