package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStsCredentialProvidersResponse Response Object
type ListStsCredentialProvidersResponse struct {
	CredentialProviders *[]StsCredentialProviderSummary `json:"credential_providers,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStsCredentialProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStsCredentialProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListStsCredentialProvidersResponse", string(data)}, " ")
}
