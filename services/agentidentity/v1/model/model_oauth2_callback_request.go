package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Oauth2CallbackRequest Request Object
type Oauth2CallbackRequest struct {

	// Unique identifier of the credential provider
	CredentialProviderId string `json:"credential_provider_id"`

	// OAuth2.0 Standard Authorization Code - one-time use, short-lived token for access token exchange. Present ONLY on successful authorization.
	Code *string `json:"code,omitempty"`

	// OAuth2.0 Standard CSRF Protection State - opaque string, echo of original request state. PRESENT FOR BOTH SUCCESS AND ERROR.
	State string `json:"state"`

	// OAuth2.0 Standard Error Code - present ONLY on authorization failure (denial/expiry/invalid). e.g. access_denied, invalid_scope, server_error
	Error *string `json:"error,omitempty"`

	// OAuth2.0 Standard Error Description - human-readable error message, paired with error param, URL-encoded for safe transmission
	ErrorDescription *string `json:"error_description,omitempty"`
}

func (o Oauth2CallbackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Oauth2CallbackRequest struct{}"
	}

	return strings.Join([]string{"Oauth2CallbackRequest", string(data)}, " ")
}
