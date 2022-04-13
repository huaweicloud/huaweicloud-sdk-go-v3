package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDomainsResponse struct {
	// 域名总数

	Total *int32 `json:"total,omitempty"`
	// 一级域名总数

	TopLevelDomainNum *int32 `json:"top_level_domain_num,omitempty"`
	// 域名列表

	Domains        *[]DomainItem `json:"domains,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainsResponse", string(data)}, " ")
}
