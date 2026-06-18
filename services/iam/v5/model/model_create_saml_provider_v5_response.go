package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSamlProviderV5Response Response Object
type CreateSamlProviderV5Response struct {
	SamlProvider   *InlineResponse201SamlProvider `json:"saml_provider,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CreateSamlProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSamlProviderV5Response struct{}"
	}

	return strings.Join([]string{"CreateSamlProviderV5Response", string(data)}, " ")
}
