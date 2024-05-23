package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInstancesReq struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”字段。该字段为false或者未提供该参数时，该条件不生效，即返回所有资源或按\"tags\"，\"matches\"等条件过滤。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含20个key，每个key下面的value最多20个，每个key对应的value可以为空数组但结构体不能缺失。key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]Tag `json:"tags,omitempty"`

	// 搜索字段,key为要匹配的字段，如resource_name等。value为匹配的值。key为固定字典值，不能包含重复的key或不支持的key。根据key的值确认是否需要模糊匹配，如resource_name默认为精确搜索（推荐实现前缀搜索），如果value为空字符串精确匹配（多数服务不存在资源名称为空的情况，因此此类情况返回空列表）。resource_id为精确匹配。第一期只做resource_name，后续再扩展。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ResourceInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstancesReq struct{}"
	}

	return strings.Join([]string{"ResourceInstancesReq", string(data)}, " ")
}
