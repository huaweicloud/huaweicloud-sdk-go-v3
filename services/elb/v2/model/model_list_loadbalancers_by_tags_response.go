package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoadbalancersByTagsResponse Response Object
type ListLoadbalancersByTagsResponse struct {

	// 总记录数。当resources为空时，表示名称为matches字段中指定的value的负载均衡器个数；resources不为空时，表示和tags字段匹配的负载均衡器的个数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 根据tag查询出的负载均衡器对象。 当请求中的action为filters，返回体中有该字段。 当请求中的action为count时，返回体中无该字段。
	Resources      *[]ResourcesByTag `json:"resources,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListLoadbalancersByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersByTagsResponse", string(data)}, " ")
}
