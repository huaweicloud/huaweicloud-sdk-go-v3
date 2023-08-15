package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tags struct {

	// 包含标签，最多包含20个key，每个key下面的value最多10个，每个key对应的value可以为空数组但结构体不能缺失。Key不能重复，同一个key中values不能重复。结果返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。如tags_any和tags字段同时存在，则去重后返回两者并集。无tag过滤条件时返回全量数据。
	Tags *[]Tag `json:"tags,omitempty"`

	// 包含任意标签，最多包含20个key，每个key下面的value最多10个。Key不能重复，同一个key中values不能重复。结果返回包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。如tags_any和tags字段同时存在，则去重后返回两者并集。无过滤条件时返回全量数据。
	TagsAny *[]Tag `json:"tags_any,omitempty"`

	// 不包含标签，最多包含20个key，每个key下面的value最多10个。Key不能重复，同一个key中values不能重复。结果返回不包含标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTags *[]Tag `json:"not_tags,omitempty"`

	// 不包含任意标签，最多包含20个key，每个key下面的value最多10个。Key不能重复，同一个key中values不能重复。结果返回不包含标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTagsAny *[]Tag `json:"not_tags_any,omitempty"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量（action为count时无此参数）从第一条数据偏移offset条数据后开始查询，如果action为filter默认为0（偏移0条数据，表示从第一条数据开始查询）,必须为数字，不能为负数
	Offset *string `json:"offset,omitempty"`

	// 操作标识（仅限于filter，count）：filter（过滤），count(查询总条数) 如果是filter就按照过滤条件查询，如果是count，只需要返回总条数，禁止返回其他字段。
	Action string `json:"action"`

	// 搜索字段,key为要匹配的字段，如resource_name等。value为匹配的值。key为固定字典值，不能包含重复的key或不支持的key。 根据key的值确认是否需要模糊匹配，如resource_name默认为模糊搜索（不区分大小写，不支持*，支持字符串匹配），如果value为空字符串则返回空列表（IEF服务不存在资源名称为空的情况，因此这类情况返回空列表）。
	Matches *[]Matches `json:"matches,omitempty"`

	// 确认是否请求源是否为console，通过该字段来判断是否返回resource_detail内容，如果为true则返回，如果为false或者不带该参数则返回空列表。
	Console *bool `json:"console,omitempty"`

	Sorted *Sorted `json:"sorted,omitempty"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
