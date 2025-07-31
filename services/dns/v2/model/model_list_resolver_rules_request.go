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

	// **参数解释：** 分页查询时配置每页返回的资源个数。 **约束限制：** 不涉及。 **取值范围：** 0~500。 **默认取值：** 500
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。 **约束限制：** 当设置marker不为空时，以marker为分页起始标识，offset不生效。。 **取值范围：** 0~2147483647。 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 分页查询的起始资源ID。 - 查询第一页时，设置为空。 - 查询下一页时，设置为上一页最后一条资源的ID。  **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Marker *string `json:"marker,omitempty"`
}

func (o ListResolverRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolverRulesRequest struct{}"
	}

	return strings.Join([]string{"ListResolverRulesRequest", string(data)}, " ")
}
