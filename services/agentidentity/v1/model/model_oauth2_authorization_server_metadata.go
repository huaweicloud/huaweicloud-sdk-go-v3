package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2AuthorizationServerMetadata Authorization server metadata for an OAuth2 provider.
type Oauth2AuthorizationServerMetadata struct {

	// Authorization endpoint of the authorization server.
	AuthorizationEndpoint string `json:"authorization_endpoint"`

	// Issuer identifier of the authorization server.
	Issuer string `json:"issuer"`

	// Token endpoint of the authorization server.
	TokenEndpoint string `json:"token_endpoint"`

	// Supported response types.
	ResponseTypes *[]string `json:"response_types,omitempty"`

	// Client authentication methods supported by the token endpoint.
	TokenEndpointAuthMethods *[]string `json:"token_endpoint_auth_methods,omitempty"`
}

func (o Oauth2AuthorizationServerMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2AuthorizationServerMetadata struct{}"
	}

	return strings.Join([]string{"Oauth2AuthorizationServerMetadata", string(data)}, " ")
}
