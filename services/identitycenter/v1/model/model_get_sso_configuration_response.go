package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetSsoConfigurationResponse Response Object
type GetSsoConfigurationResponse struct {
	SsoConfiguration *SsoConfigurationDto `json:"sso_configuration,omitempty"`
	HttpStatusCode   int                  `json:"-"`
}

func (o GetSsoConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetSsoConfigurationResponse struct{}"
	}

	return strings.Join([]string{"GetSsoConfigurationResponse", string(data)}, " ")
}
