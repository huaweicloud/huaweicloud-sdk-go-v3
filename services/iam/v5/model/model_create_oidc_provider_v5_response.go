package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOidcProviderV5Response Response Object
type CreateOidcProviderV5Response struct {
	OidcProvider   *InlineResponse2011OidcProvider `json:"oidc_provider,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateOidcProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOidcProviderV5Response struct{}"
	}

	return strings.Join([]string{"CreateOidcProviderV5Response", string(data)}, " ")
}
