package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomJwtAuthorizerConfiguration struct {

	// This URL is used to fetch OpenID Connect configuration.
	DiscoveryUrl string `json:"discovery_url"`

	AllowedAudience *[]string `json:"allowed_audience,omitempty"`

	AllowedClients *[]string `json:"allowed_clients,omitempty"`

	AllowedScopes *[]string `json:"allowed_scopes,omitempty"`

	// Custom claim validation rules applied to inbound JWTs.
	CustomClaims *[]CustomClaimValidation `json:"custom_claims,omitempty"`
}

func (o CustomJwtAuthorizerConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomJwtAuthorizerConfiguration struct{}"
	}

	return strings.Join([]string{"CustomJwtAuthorizerConfiguration", string(data)}, " ")
}
