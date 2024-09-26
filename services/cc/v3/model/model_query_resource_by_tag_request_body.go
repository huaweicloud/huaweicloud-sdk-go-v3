package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResourceByTagRequestBody struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”、“tags_any”、“not_tags”、“not_tags_any”字段。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 包含标签，最多包含20个key，每个key下面的value最多10个，每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]QueryTag `json:"tags,omitempty"`

	// 包含任意标签，最多包含20个key，每个key下面的value最多10个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	TagsAny *[]QueryTag `json:"tags_any,omitempty"`

	// 不包含标签，最多包含20个key，每个key下面的value最多10个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回不包含标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTags *[]QueryTag `json:"not_tags,omitempty"`

	// 不包含任意标签，最多包含20个key，每个key下面的value最多10个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回不包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTagsAny *[]QueryTag `json:"not_tags_any,omitempty"`

	// 是否匹配以下tag，key必须为\"resource_name\"，value如果有值则模糊匹配，如果为空字符串则精确匹配。
	Matches *[]TmsMatch `json:"matches,omitempty"`
}

func (o QueryResourceByTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourceByTagRequestBody struct{}"
	}

	return strings.Join([]string{"QueryResourceByTagRequestBody", string(data)}, " ")
}
