package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowConfigurationParameterRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationParameterRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationParameterRequest", string(data)}, " ")
}
