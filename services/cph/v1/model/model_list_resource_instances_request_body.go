package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceInstancesRequestBody 查询资源实例请求体
type ListResourceInstancesRequestBody struct {

	// 操作标识（仅限于filter，count）：filter（过滤），count(查询总条数)。 为filter时表示分页查询，为count只需按照条件将总条数返回即可。
	Action ListResourceInstancesRequestBodyAction `json:"action"`

	// 索引位置， 从offset指定的下一条数据开始查询。 查询第一页数据时，不需要传入此参数，查询后续页码数据时，将查询前一页数据时响应体中的值带入此参数。  action为count时无此参数。  action为filter时，默认为0，必须为数字，且不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数。  action为count时无此参数。  action为filter时，默认为1000。limit最多为1000，不能为负数，最小值为1。
	Limit *string `json:"limit,omitempty"`

	// 不包含任意一个标签，该字段为true时查询所有不带标签的资源。
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	// 最多包含10个key，每个key最多包含10个value，结构体不能缺失。key不能为空或者空字符串。key不能重复，同一个key中value不能重复，不同key对应的资源之间为与的关系。
	Tags *[]Tags `json:"tags,omitempty"`

	// 搜索字段。  key为要匹配的字段，当前只支持resource_name。  value为匹配的值，当前为精确匹配。
	Matches *[]Match `json:"matches,omitempty"`
}

func (o ListResourceInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequestBody", string(data)}, " ")
}

type ListResourceInstancesRequestBodyAction struct {
	value string
}

type ListResourceInstancesRequestBodyActionEnum struct {
	FILTER ListResourceInstancesRequestBodyAction
	COUNT  ListResourceInstancesRequestBodyAction
}

func GetListResourceInstancesRequestBodyActionEnum() ListResourceInstancesRequestBodyActionEnum {
	return ListResourceInstancesRequestBodyActionEnum{
		FILTER: ListResourceInstancesRequestBodyAction{
			value: "filter",
		},
		COUNT: ListResourceInstancesRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListResourceInstancesRequestBodyAction) Value() string {
	return c.value
}

func (c ListResourceInstancesRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceInstancesRequestBodyAction) UnmarshalJSON(b []byte) error {
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
