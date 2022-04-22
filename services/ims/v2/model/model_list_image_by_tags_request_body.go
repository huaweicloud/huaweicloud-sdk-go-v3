package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 请求参数
type ListImageByTagsRequestBody struct {

	// 操作标识（区分大小写），支持filter、count。filter就是分页查询；count是只需按照条件将总条数返回即可。
	Action ListImageByTagsRequestBodyAction `json:"action"`

	// 包含标签，最多包含10个key，每个key对应的values最多包含10个值，且key和values都不能重复。不能为空列表。
	Tags *[]Tags `json:"tags,omitempty"`

	// 包含任意标签，最多包含10个key，每个key对应的values最多包含10个值，且key和values都不能重复。不允许为空列表，但可以不传递参数。
	TagsAny *[]Tags `json:"tags_any,omitempty"`

	// 不包含标签，最多包含10个key，每个key对应的values最多包含10个值，且key和values都不能重复。不能为空列表。
	NotTags *[]Tags `json:"not_tags,omitempty"`

	// 不包含任意标签，最多包含10个key，每个key对应的values最多包含10个值，且key和values都不能重复。不能为空列表。
	NotTagsAny *[]Tags `json:"not_tags_any,omitempty"`

	// 最大查询记录数(action为count，时此参数无效）如果action为filter默认为10，limit最多为1000，不能为负数，最小值为1。
	Limit *string `json:"limit,omitempty"`

	// 索引位置， 从offset指定的下一条数据开始查询。 查询第一页数据时，不需要传入此参数（action为count时，此参数无效），如果action为filter默认为0，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 搜索字段，key为要匹配的字段，如resource_name、resource_id等。value为匹配的值。多个matches的key不允许重复。不允许为空列表，但可以不传递参数。
	Matches *[]TagKeyValue `json:"matches,omitempty"`

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源，此时忽略tag、not_tags、tags_any、not_tags_any字段。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
}

func (o ListImageByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListImageByTagsRequestBody", string(data)}, " ")
}

type ListImageByTagsRequestBodyAction struct {
	value string
}

type ListImageByTagsRequestBodyActionEnum struct {
	FILTER ListImageByTagsRequestBodyAction
	COUNT  ListImageByTagsRequestBodyAction
}

func GetListImageByTagsRequestBodyActionEnum() ListImageByTagsRequestBodyActionEnum {
	return ListImageByTagsRequestBodyActionEnum{
		FILTER: ListImageByTagsRequestBodyAction{
			value: "filter",
		},
		COUNT: ListImageByTagsRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListImageByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListImageByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
