package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventResultSort 返回列表的排序方式，可以为空
type EventResultSort struct {

	// 排序字段列表。会根据列表中定义顺序对返回列表最排序
	OrderBy *[]string `json:"order_by,omitempty"`

	// 排序方式枚举值。asc代表正序，desc代表倒序
	Order *EventResultSortOrder `json:"order,omitempty"`
}

func (o EventResultSort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResultSort struct{}"
	}

	return strings.Join([]string{"EventResultSort", string(data)}, " ")
}

type EventResultSortOrder struct {
	value string
}

type EventResultSortOrderEnum struct {
	ASC  EventResultSortOrder
	DESC EventResultSortOrder
}

func GetEventResultSortOrderEnum() EventResultSortOrderEnum {
	return EventResultSortOrderEnum{
		ASC: EventResultSortOrder{
			value: "asc",
		},
		DESC: EventResultSortOrder{
			value: "desc",
		},
	}
}

func (c EventResultSortOrder) Value() string {
	return c.value
}

func (c EventResultSortOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventResultSortOrder) UnmarshalJSON(b []byte) error {
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
