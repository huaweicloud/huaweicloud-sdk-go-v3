package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourcesByTagsResponse Response Object
type ListResourcesByTagsResponse struct {

	// 操作标识(仅限于filter，count)  - filter：分页查询 - count：查询总条数，只需按照条件将总条数返回即可
	Action *ListResourcesByTagsResponseAction `json:"action,omitempty"`

	// 查询记录数(action为count时无此参数)如果action为filter默认为1000，limit最多为1000，不能为负数，最小值为1
	Limit *string `json:"limit,omitempty"`

	// 索引位置, 从offset指定的下一条数据开始查询。 查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数(action为count时无此参数)如果action为filter默认为0，必须为数字，不能为负数
	Offset *string `json:"offset,omitempty"`

	// 返回结果包含该参数中所有标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。
	Tags *[]Tags `json:"tags,omitempty"`

	// 返回结果包含该参数中任意一个标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。
	TagsAny *[]Tags `json:"tags_any,omitempty"`

	// 返回结果不包含该参数中所有标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个, 结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。
	NotTags *[]Tags `json:"not_tags,omitempty"`

	// 返回结果不包含该参数中任意一个标签对应的资源，该参数最多包含10个key，每个key下面的value最多10个，结构体不能缺失，key不能为空或者空字符串。Key不能重复，同一个key中values不能重复。
	NotTagsAny *[]Tags `json:"not_tags_any,omitempty"`

	// 搜索字段，key为要匹配的字段，当前仅支持resource_name。value为匹配的值。此字段为固定字典值
	Matches        *string `json:"matches,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourcesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsResponse", string(data)}, " ")
}

type ListResourcesByTagsResponseAction struct {
	value string
}

type ListResourcesByTagsResponseActionEnum struct {
	FILTER ListResourcesByTagsResponseAction
	COUNT  ListResourcesByTagsResponseAction
}

func GetListResourcesByTagsResponseActionEnum() ListResourcesByTagsResponseActionEnum {
	return ListResourcesByTagsResponseActionEnum{
		FILTER: ListResourcesByTagsResponseAction{
			value: "filter",
		},
		COUNT: ListResourcesByTagsResponseAction{
			value: "count",
		},
	}
}

func (c ListResourcesByTagsResponseAction) Value() string {
	return c.value
}

func (c ListResourcesByTagsResponseAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourcesByTagsResponseAction) UnmarshalJSON(b []byte) error {
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
