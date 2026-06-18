package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSamlProvidersV5Response Response Object
type ListSamlProvidersV5Response struct {

	// **参数解释**： SAML 提供商。  **取值范围**： 不涉及。
	SamlProviders *[]InlineResponse200SamlProviders `json:"saml_providers,omitempty"`

	PageInfo       *InlineResponse200PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSamlProvidersV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSamlProvidersV5Response struct{}"
	}

	return strings.Join([]string{"ListSamlProvidersV5Response", string(data)}, " ")
}
