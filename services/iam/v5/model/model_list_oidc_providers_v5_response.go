package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOidcProvidersV5Response Response Object
type ListOidcProvidersV5Response struct {

	// **参数解释**： OIDC 提供商列表。  **取值范围**： 不涉及。
	OidcProviders *[]InlineResponse2002OidcProviders `json:"oidc_providers,omitempty"`

	PageInfo       *InlineResponse200PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListOidcProvidersV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOidcProvidersV5Response struct{}"
	}

	return strings.Join([]string{"ListOidcProvidersV5Response", string(data)}, " ")
}
