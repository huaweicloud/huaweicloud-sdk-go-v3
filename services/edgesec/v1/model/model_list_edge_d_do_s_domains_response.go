package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeDDoSDomainsResponse Response Object
type ListEdgeDDoSDomainsResponse struct {

	// 域名列表
	DomainList *[]EdgeDDoSDomainVo `json:"domain_list,omitempty"`

	// 域名总条目
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEdgeDDoSDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeDDoSDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeDDoSDomainsResponse", string(data)}, " ")
}
