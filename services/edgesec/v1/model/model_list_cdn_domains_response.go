package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCdnDomainsResponse Response Object
type ListCdnDomainsResponse struct {

	// 全部CDN域名的数量
	Total *int32 `json:"total,omitempty"`

	// 查询结果CDN域名的数量
	Count *int32 `json:"count,omitempty"`

	// 详细的CDN域名信息
	Domains        *[]ShowCdnDomainResponseBody `json:"domains,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListCdnDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCdnDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListCdnDomainsResponse", string(data)}, " ")
}
