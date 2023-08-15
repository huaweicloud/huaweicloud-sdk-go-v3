package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSharesByTagsReqBody ResourceInstance操作的请求体。
type ResourceSharesByTagsReqBody struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含20个key，每个key下面的value最多20个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回包含所有标签的资源列表，key之间是\"与\"的关系，key-value结构中value是\"或\"的关系。无tag过滤条件时返回全量数据。
	Tags *[]TagFilter `json:"tags,omitempty"`

	// 资源共享实例绑定标签列表的匹配字段。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ResourceSharesByTagsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSharesByTagsReqBody struct{}"
	}

	return strings.Join([]string{"ResourceSharesByTagsReqBody", string(data)}, " ")
}
