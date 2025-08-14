package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetSpConfigurationForDirectoryResponse Response Object
type GetSpConfigurationForDirectoryResponse struct {
	SpOidcConfig *SpoidcConfig `json:"sp_oidc_config,omitempty"`

	SpSamlConfig   *SpsamlConfig `json:"sp_saml_config,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o GetSpConfigurationForDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetSpConfigurationForDirectoryResponse struct{}"
	}

	return strings.Join([]string{"GetSpConfigurationForDirectoryResponse", string(data)}, " ")
}
