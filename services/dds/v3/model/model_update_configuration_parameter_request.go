package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigurationParameterRequest Request Object
type UpdateConfigurationParameterRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	Body *UpdateConfigurationParameterRequestBody `json:"body,omitempty"`
}

func (o UpdateConfigurationParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterRequest", string(data)}, " ")
}
