package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListInstancesByTagsRequestBody struct {

	// **参数解释：** 索引位置偏移量，表示从第一条数据偏移offset条数据后开始查询。 **约束限制：**   - “action”值为“count”时，不传该参数。   - “action”值为“filter”时，取值必须为数字，不能为负数。 **取值范围：** 不涉及。 **默认取值：** 默认取0值，表示从第一条数据开始查询。
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** 查询记录数。 **约束限制：**   - “action”值为“count”时，不传该参数。   - “action”值为“filter”时，取值范围：1~100。不传该参数时，默认查询前100条实例信息。 **取值范围：** 1~100 **默认取值：** 不涉及。
	Limit *string `json:"limit,omitempty"`

	// **参数解释：** 操作标识。 **约束限制：** - 取值为“count”，表示仅返回总记录数，禁止返回其他字段。 - 取值为“filter”，表示根据标签过滤条件查询实例。 **取值范围：** count和filter **默认取值：** 不涉及。
	Action ListInstancesByTagsRequestBodyAction `json:"action"`

	// **参数解释：** 搜索字段。 **约束限制：**   - 该字段值为空，表示不按照实例名称或实例ID查询。   - 该字段值不为空， **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Matches *[]QueryMatchItem `json:"matches,omitempty"`

	// **参数解释：** 包含标签。 **约束限制：** 最多包含10个key。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Tags *[]QueryTagItem `json:"tags,omitempty"`
}

func (o ListInstancesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListInstancesByTagsRequestBody", string(data)}, " ")
}

type ListInstancesByTagsRequestBodyAction struct {
	value string
}

type ListInstancesByTagsRequestBodyActionEnum struct {
	COUNT  ListInstancesByTagsRequestBodyAction
	FILTER ListInstancesByTagsRequestBodyAction
}

func GetListInstancesByTagsRequestBodyActionEnum() ListInstancesByTagsRequestBodyActionEnum {
	return ListInstancesByTagsRequestBodyActionEnum{
		COUNT: ListInstancesByTagsRequestBodyAction{
			value: "count",
		},
		FILTER: ListInstancesByTagsRequestBodyAction{
			value: "filter",
		},
	}
}

func (c ListInstancesByTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ListInstancesByTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesByTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
