package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainsResponse Response Object
type ListDomainsResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	// 域名列表。
	Items *[]DomainItem `json:"items,omitempty"`

	// API类型，固定值“Domain”，该值不可修改。
	Kind           *string `json:"kind,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainsResponse", string(data)}, " ")
}
