package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询资源实例的请求体。
type ListTagResourceInstancesRequestBody struct {

	// 索引位置， 从offset指定的下一条数据开始查询。 查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数（action为count时无此参数）如果action为filter默认为0,必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1。
	Limit *string `json:"limit,omitempty"`

	// 操作标识（仅限于filter，count）：filter（过滤），count(查询总条数) 如果是filter就是分页查询，如果是count只需按照条件将总条数返回即可。
	Action ListTagResourceInstancesRequestBodyAction `json:"action"`

	// 搜索字段,key为要匹配的字段，如resource_name等。value为匹配的值。此字段为固定字典值。 根据不同的字段确认是否需要模糊匹配，如resource_name默认为模糊搜索（不区分大小写），如果value为空字符串精确匹配。resource_id为精确匹配。第一期只做resource_name，后续在扩展。
	Matches *[]Match `json:"matches,omitempty"`

	// 不包含标签，最多包含10个key，每个key下面的value最多10个, 结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回不包含标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTags *[]Tags `json:"not_tags,omitempty"`

	// 包含标签，最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回包含所有标签的资源列表，key之间是与的关系，key-value结构中value是或的关系。无tag过滤条件时返回全量数据。
	Tags *[]Tags `json:"tags,omitempty"`

	// 包含任意标签，最多包含10个key，每个key下面的value最多10个,结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回包含任意标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	TagsAny *[]Tags `json:"tags_any,omitempty"`

	// 不包含任意标签，最多包含10个key，每个key下面的value最多10个,结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。返回不包含任意标签的资源列表，key之间是或的关系，key-value结构中value是或的关系。无过滤条件时返回全量数据。
	NotTagsAny *[]Tags `json:"not_tags_any,omitempty"`
}

func (o ListTagResourceInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagResourceInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListTagResourceInstancesRequestBody", string(data)}, " ")
}

type ListTagResourceInstancesRequestBodyAction struct {
	value string
}

type ListTagResourceInstancesRequestBodyActionEnum struct {
	FILTER ListTagResourceInstancesRequestBodyAction
	COUNT  ListTagResourceInstancesRequestBodyAction
}

func GetListTagResourceInstancesRequestBodyActionEnum() ListTagResourceInstancesRequestBodyActionEnum {
	return ListTagResourceInstancesRequestBodyActionEnum{
		FILTER: ListTagResourceInstancesRequestBodyAction{
			value: "filter",
		},
		COUNT: ListTagResourceInstancesRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListTagResourceInstancesRequestBodyAction) Value() string {
	return c.value
}

func (c ListTagResourceInstancesRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTagResourceInstancesRequestBodyAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
