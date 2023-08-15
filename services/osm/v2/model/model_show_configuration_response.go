package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigurationResponse Response Object
type ShowConfigurationResponse struct {

	// 配置项键
	ConfigKey *string `json:"config_key,omitempty"`

	// 配置项值
	ConfigValue *string `json:"config_value,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationResponse", string(data)}, " ")
}
