package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSamlProviderV5Response Response Object
type ShowSamlProviderV5Response struct {
	SamlProvider   *InlineResponse2001SamlProvider `json:"saml_provider,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowSamlProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSamlProviderV5Response struct{}"
	}

	return strings.Join([]string{"ShowSamlProviderV5Response", string(data)}, " ")
}
