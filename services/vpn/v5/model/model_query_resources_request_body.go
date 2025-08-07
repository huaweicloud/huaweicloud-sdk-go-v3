package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResourcesRequestBody struct {

	// 该字段为true时查询所有不带标签的资源，此时忽略 “tags”字段。该字段为false或者未提供该参数时，该条件不生效，即返回所有资源或按\"tags\"，\"matches\"等条件过滤
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 标签列表
	Tags *[]Tag `json:"tags,omitempty"`

	// 搜索字段，包含key和value。key为要匹配的字段，如resource_name等。value为匹配的值。key为固定字典值
	Matches *[]Match `json:"matches,omitempty"`
}

func (o QueryResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"QueryResourcesRequestBody", string(data)}, " ")
}
