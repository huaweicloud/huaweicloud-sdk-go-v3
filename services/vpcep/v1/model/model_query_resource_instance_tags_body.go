package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryResourceInstanceTagsBody 查询资源实例接口请求结构体
type QueryResourceInstanceTagsBody struct {

	// 包含标签，最多包含10个key，每个key下面的value最多10个， 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复， 同一个key中values不能重复。结果返回包含所有标签的资源列表， key之间是与的关系，key-value结构中value是或的关系。 无tag过滤条件时返回全量数据。
	Tags *[]TagValuesList `json:"tags,omitempty"`

	// 包含任意标签，最多包含10个key，每个key下面的value最多10个， 每个key对应的value可以为空数组但结构体不能缺失。 Key不能重复，同一个key中values不能重复。 结果返回包含标签的资源列表，key之间是或的关系，key、value结构中value是或的关系。 无过滤条件时返回全量数据。
	TagsAny *[]TagValuesList `json:"tags_any,omitempty"`

	// 不包含标签，最多包含10个key，每个key下面的value最多10个， 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复， 同一个key中values不能重复。结果返回不包含标签的资源列表，key之间是与的关系， key、value结构中value是或的关系。 无过滤条件时返回全量数据。
	NotTags *[]TagValuesList `json:"not_tags,omitempty"`

	// 不包含任意标签，最多包含10个key，每个key下面的value最多10个， 每个key对应的value可以为空数组但结构体不能缺失。Key不能重复， 同一个key中values不能重复。结果返回不包含标签的资源列表， key之间是与的关系，key、value结构中value是或的关系。 无过滤条件时返回全量数据。
	NotTagsAny *[]TagValuesList `json:"not_tags_any,omitempty"`

	// 系统标签，
	SysTags *[]TagValuesList `json:"sys_tags,omitempty"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000， limit最多为1000，不能为负数，最小值为1。
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量（action为count时无此参数）从第一条数据偏移offset条数据后开始查询， 如果action为filter默认为0（偏移0条数据，表示从第一条数据开始查询）， 必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 操作标识（仅限于filter，count）：filter（过滤）， count(查询总条数)如果是filter就按照过滤条件查询， 如果是count，只需要返回总条数，禁止返回其他字段。
	Action QueryResourceInstanceTagsBodyAction `json:"action"`

	// 搜索字段，key为要匹配的字段，如resource_name等。value为匹配的值。 key为固定字典值，不能包含重复的key或不支持的key。 根据key的值确认是否需要模糊匹配，如resource_name默认为模糊搜索（不区分大小写）， 如果value为空字符串精确匹配（多数服务不存在资源名称为空的情况， 因此此类情况返回空列表）。resource_id为精确匹配。 第一期只做resource_name，后续再扩展。
	Matches *[]Match `json:"matches,omitempty"`

	// 默认为false，取值【true/false】,当withoutAnyTag=true， 忽略tags、tagsAny、notTags、notTagsAny参数校验。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
}

func (o QueryResourceInstanceTagsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourceInstanceTagsBody struct{}"
	}

	return strings.Join([]string{"QueryResourceInstanceTagsBody", string(data)}, " ")
}

type QueryResourceInstanceTagsBodyAction struct {
	value string
}

type QueryResourceInstanceTagsBodyActionEnum struct {
	FILTER QueryResourceInstanceTagsBodyAction
	COUNT  QueryResourceInstanceTagsBodyAction
}

func GetQueryResourceInstanceTagsBodyActionEnum() QueryResourceInstanceTagsBodyActionEnum {
	return QueryResourceInstanceTagsBodyActionEnum{
		FILTER: QueryResourceInstanceTagsBodyAction{
			value: "filter",
		},
		COUNT: QueryResourceInstanceTagsBodyAction{
			value: "count",
		},
	}
}

func (c QueryResourceInstanceTagsBodyAction) Value() string {
	return c.value
}

func (c QueryResourceInstanceTagsBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryResourceInstanceTagsBodyAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
