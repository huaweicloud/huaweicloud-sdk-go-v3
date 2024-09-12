package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainsResponse Response Object
type ShowDomainsResponse struct {

	// 全部防护域名的数量
	Total *int64 `json:"total,omitempty"`

	// 域名已使用配额
	DomainUsedQuota *int64 `json:"domain_used_quota,omitempty"`

	// 详细的防护域名信息
	DomainList     *[]DomainInfo `json:"domain_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainsResponse", string(data)}, " ")
}
