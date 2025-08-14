package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceServiceProviderConfigurationResponse Response Object
type UpdateApplicationInstanceServiceProviderConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApplicationInstanceServiceProviderConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceServiceProviderConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceServiceProviderConfigurationResponse", string(data)}, " ")
}
