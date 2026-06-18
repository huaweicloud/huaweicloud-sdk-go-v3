package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOidcProviderV5Response Response Object
type ShowOidcProviderV5Response struct {
	OidcProvider   *InlineResponse2003OidcProvider `json:"oidc_provider,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowOidcProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOidcProviderV5Response struct{}"
	}

	return strings.Join([]string{"ShowOidcProviderV5Response", string(data)}, " ")
}
