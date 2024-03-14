package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceInstanceReqBody ResourceInstance操作的请求体。
type ResourceInstanceReqBody struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]TagsDto `json:"tags,omitempty"`

	// 要绑定到新创建的账号的标签列表。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ResourceInstanceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceReqBody struct{}"
	}

	return strings.Join([]string{"ResourceInstanceReqBody", string(data)}, " ")
}
