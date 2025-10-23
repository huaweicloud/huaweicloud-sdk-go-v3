package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetProtectionConfigurationResponse Response Object
type SetProtectionConfigurationResponse struct {

	// 配置保护的返回体
	Body           *[]ProtectionConfigurationResponseItem `json:"body,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o SetProtectionConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetProtectionConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SetProtectionConfigurationResponse", string(data)}, " ")
}
