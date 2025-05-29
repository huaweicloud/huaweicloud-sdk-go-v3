package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCertificatesByTagRequestBody struct {

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略 “tags”、“tags_any”、“not_tags”、“not_tags_any”字段。  - true  - false
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 标签列表。 包含标签，最多包含20个key，每个key下面的value最多20个，每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]ScsTag `json:"tags,omitempty"`

	// 标签列表。 包含任意标签，最多包含20个key，每个key下面的value最多20个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	TagsAny *[]ScsTag `json:"tags_any,omitempty"`

	// 标签列表。 不包含标签，最多包含20个key，每个key下面的value最多20个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回不包含标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTags *[]ScsTag `json:"not_tags,omitempty"`

	// 标签列表。 不包含任意标签，最多包含20个key，每个key下面的value最多20个, 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回不包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTagsAny *[]ScsTag `json:"not_tags_any,omitempty"`

	// 每页条目数量，取值如下： - 10：每页显示10条证书信息。 - 20：每页显示20条证书信息。 - 50：每页显示50条证书信息。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量，从offset指定的下一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 操作标识（可设置为“filter”或者“count”）。  - filter：表示过滤。  - count：表示查询总条数。
	Action *string `json:"action,omitempty"`

	// 搜索字段。 key为要匹配的字段，如resource_name等。value为匹配的值。key为固定字典值，不能包含重复的key或不支持的key。
	Matches *[]ScsMatch `json:"matches,omitempty"`
}

func (o ListCertificatesByTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesByTagRequestBody struct{}"
	}

	return strings.Join([]string{"ListCertificatesByTagRequestBody", string(data)}, " ")
}
