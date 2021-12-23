package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListResourceReq struct {
	// 返回结果包含该参数中所有标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。

	Tags *[]TagWithMultiValue `json:"tags,omitempty"`
	// 返回结果包含该参数中任意一个标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。

	TagsAny *[]TagWithMultiValue `json:"tags_any,omitempty"`
	// 返回结果不包含该参数中所有标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个, 结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。

	NotTags *[]TagWithMultiValue `json:"not_tags,omitempty"`
	// 返回结果不包含该参数中任意一个标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。

	NotTagsAny *[]TagWithMultiValue `json:"not_tags_any,omitempty"`
	// 操作标识（仅限于filter，count）：filter（过滤），count(查询总条数)。  如果是filter则为分页查询，如果是count会按照条件将总条数返回。

	Action ListResourceReqAction `json:"action"`
	// 查询记录数（action为count时无此参数）如果action为filter默认为1000，limit最多为1000,不能为负数，最小值为1。

	Limit *int32 `json:"limit,omitempty"`
	// （索引位置），从offset指定的下一条数据开始查询。查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数（action为count时无此参数）如果action为filter默认为0,必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 搜索字段，key为要匹配的字段，如resource_name等。value为匹配的值。此字段为固定字典值。 根据不同的字段确认是否需要模糊匹配，如resource_name默认为模糊搜索，如果value为空字符串精确匹配。

	Matches *[]Match `json:"matches,omitempty"`
}

func (o ListResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceReq struct{}"
	}

	return strings.Join([]string{"ListResourceReq", string(data)}, " ")
}

type ListResourceReqAction struct {
	value string
}

type ListResourceReqActionEnum struct {
	FILTER ListResourceReqAction
	COUNT  ListResourceReqAction
}

func GetListResourceReqActionEnum() ListResourceReqActionEnum {
	return ListResourceReqActionEnum{
		FILTER: ListResourceReqAction{
			value: "filter",
		},
		COUNT: ListResourceReqAction{
			value: "count",
		},
	}
}

func (c ListResourceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceReqAction) UnmarshalJSON(b []byte) error {
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
