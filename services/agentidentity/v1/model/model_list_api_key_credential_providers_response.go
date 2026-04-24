package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiKeyCredentialProvidersResponse Response Object
type ListApiKeyCredentialProvidersResponse struct {
	CredentialProviders *[]ApiKeyCredentialProviderSummary `json:"credential_providers,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListApiKeyCredentialProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiKeyCredentialProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListApiKeyCredentialProvidersResponse", string(data)}, " ")
}
