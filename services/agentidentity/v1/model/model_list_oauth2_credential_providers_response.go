package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOauth2CredentialProvidersResponse Response Object
type ListOauth2CredentialProvidersResponse struct {
	CredentialProviders *[]Oauth2CredentialProviderSummary `json:"credential_providers,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOauth2CredentialProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOauth2CredentialProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListOauth2CredentialProvidersResponse", string(data)}, " ")
}
