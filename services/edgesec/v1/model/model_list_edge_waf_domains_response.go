package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeWafDomainsResponse Response Object
type ListEdgeWafDomainsResponse struct {

	// 全部防护域名的数量
	Total *int64 `json:"total,omitempty"`

	// 详细的防护域名信息
	DomainList     *[]ShowWafDomainResponseBody `json:"domain_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListEdgeWafDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeWafDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeWafDomainsResponse", string(data)}, " ")
}
