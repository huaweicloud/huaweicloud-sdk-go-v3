package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询资源实例接口请求结构体
type QueryResourceInstanceTagsBody struct {
	// 包含标签，最多包含10个key，每 个key下面的value最多10个，每 个key对应的value可以为空数组但 结构体不能缺失。Key不能重复， 同一个key中values不能重复。结 果返回包含所有标签的资源列表， key之间是与的关系，key-value结 构中value是或的关系。无tag过滤 条件时返回全量数据。

	Tags *[]TagValuesList `json:"tags,omitempty"`
	// 包含任意标签，最多包含10个 key，每个key下面的value最多10 个，每个key对应的value可以为空 数组但结构体不能缺失。Key不能 重复，同一个key中values不能重 复。结果返回包含标签的资源列 表，key之间是或的关系，keyvalue 结构中value是或的关系。无 过滤条件时返回全量数据。

	TagsAny *[]TagAnyValuesList `json:"tags_any,omitempty"`
	// 不包含标签，最多包含10个key， 每个key下面的value最多10个， 每个key对应的value可以为空数组 但结构体不能缺失。Key不能重 复，同一个key中values不能重 复。结果返回不包含标签的资源列 表，key之间是与的关系，keyvalue 结构中value是或的关系。无 过滤条件时返回全量数据。

	NotTags *[]NotTagsValuesList `json:"not_tags,omitempty"`
	// 不包含任意标签，最多包含10个 key，每个key下面的value最多10 个，每个key对应的value可以为空 数组但结构体不能缺失。Key不能 重复，同一个key中values不能重 复。结果返回不包含标签的资源列 表，key之间是与的关系，keyvalue 结构中value是或的关系。无 过滤条件时返回全量数据。

	NotTagsAny *[]NotTagsAnyValuesList `json:"not_tags_any,omitempty"`
	// 查询记录数（action为count时无 此参数）如果action为filter默认为 1000，limit最多为1000，不能为 负数，最小值为1。

	Limit *string `json:"limit,omitempty"`
	// 索引位置，偏移量（action为 count时无此参数）从第一条数据 偏移offset条数据后开始查询，如 果action为filter默认为0（偏移0 条数据，表示从第一条数据开始查 询），必须为数字，不能为负数。

	Offset *string `json:"offset,omitempty"`
	// 操作标识（仅限于filter， count）：filter（过滤）， count(查询总条数) 如果是filter就按照过滤条件查 询，如果是count，只需要返回总 条数，禁止返回其他字段。

	Action *string `json:"action,omitempty"`
	// 搜索字段，key为要匹配的字段， 如resource_name等。value为匹 配的值。key为固定字典值，不能 包含重复的key或不支持的key。 根据key的值确认是否需要模糊匹 配，如resource_name默认为模糊 搜索（不区分大小写），如果 value为空字符串精确匹配（多数 服务不存在资源名称为空的情况， 因此此类情况返回空列表）。 resource_id为精确匹配。第一期 只做resource_name，后续再扩 展。

	Matches *[]TagMatchList `json:"matches,omitempty"`
}

func (o QueryResourceInstanceTagsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourceInstanceTagsBody struct{}"
	}

	return strings.Join([]string{"QueryResourceInstanceTagsBody", string(data)}, " ")
}
