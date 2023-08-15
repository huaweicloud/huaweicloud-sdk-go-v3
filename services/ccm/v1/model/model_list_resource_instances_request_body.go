package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListResourceInstancesRequestBody struct {

	// 标签列表。 最多包含20个key，每个key下面的value最多20个，每个key对应的value可以为空数组但结构体不能缺失。key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]DomainTags `json:"tags,omitempty"`

	// 每页条目数量，取值如下： - 10：每页显示10条资源信息。 - 20：每页显示20条资源信息。 - 50：每页显示50条资源信息。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量，从offset指定的下一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 搜索字段。 key为要匹配的字段，如resource_name等。value为匹配的值。key为固定字典值，不能包含重复的key或不支持的key。
	Matches *[]ResourceTag `json:"matches,omitempty"`
}

func (o ListResourceInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequestBody", string(data)}, " ")
}
