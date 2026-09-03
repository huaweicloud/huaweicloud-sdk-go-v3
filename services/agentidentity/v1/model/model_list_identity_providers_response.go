package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIdentityProvidersResponse Response Object
type ListIdentityProvidersResponse struct {

	// List of identity providers.
	IdentityProviders *[]IdentityProviderListSummary `json:"identity_providers,omitempty"`
	HttpStatusCode    int                            `json:"-"`
}

func (o ListIdentityProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIdentityProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListIdentityProvidersResponse", string(data)}, " ")
}
