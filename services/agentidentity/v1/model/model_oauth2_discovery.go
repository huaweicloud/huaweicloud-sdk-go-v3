package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2Discovery Discovery information for an OAuth2 provider (UNION). Only one member can be specified.
type Oauth2Discovery struct {

	// This URL is used to fetch OpenID Connect configuration.
	DiscoveryUrl *string `json:"discovery_url,omitempty"`

	AuthorizationServerMetadata *Oauth2AuthorizationServerMetadata `json:"authorization_server_metadata,omitempty"`
}

func (o Oauth2Discovery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2Discovery struct{}"
	}

	return strings.Join([]string{"Oauth2Discovery", string(data)}, " ")
}
