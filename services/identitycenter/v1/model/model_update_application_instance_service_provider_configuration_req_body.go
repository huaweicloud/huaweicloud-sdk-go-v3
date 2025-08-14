package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceServiceProviderConfigurationReqBody UpdateApplicationInstanceServiceProviderConfiguration的请求体
type UpdateApplicationInstanceServiceProviderConfigurationReqBody struct {
	ServiceProviderConfig *ServiceProviderConfigDto `json:"service_provider_config"`
}

func (o UpdateApplicationInstanceServiceProviderConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceServiceProviderConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceServiceProviderConfigurationReqBody", string(data)}, " ")
}
