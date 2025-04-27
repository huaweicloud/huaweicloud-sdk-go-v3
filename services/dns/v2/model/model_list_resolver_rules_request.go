package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResolverRulesRequest Request Object
type ListResolverRulesRequest struct {

	// 待查询的转发规则的域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 待查询的转发规则的名称。
	Name *string `json:"name,omitempty"`

	// 终端节点ID。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 转发规则ID。
	Id *string `json:"id,omitempty"`

	// 每页返回的资源个数。  取值范围：0~500  取值一般为10，20，50。默认值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当前设置marker不为空时，以marker为分页起始标识。
	Offset *int32 `json:"offset,omitempty"`

	// 分页查询起始的资源ID，为空时为查询第一页。  默认值为空。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResolverRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolverRulesRequest struct{}"
	}

	return strings.Join([]string{"ListResolverRulesRequest", string(data)}, " ")
}
