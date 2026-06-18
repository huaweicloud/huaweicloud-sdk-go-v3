package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSamlProviderV5Response Response Object
type UpdateSamlProviderV5Response struct {
	SamlProvider   *InlineResponse201SamlProvider `json:"saml_provider,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o UpdateSamlProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSamlProviderV5Response struct{}"
	}

	return strings.Join([]string{"UpdateSamlProviderV5Response", string(data)}, " ")
}
