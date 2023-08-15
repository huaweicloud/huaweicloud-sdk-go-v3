package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OrderParam struct {

	// timeUsed：响应时间，startTime：产生时间。
	Field *string `json:"field,omitempty"`

	// ASC：正序，DESC：逆序。
	Order *OrderParamOrder `json:"order,omitempty"`
}

func (o OrderParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderParam struct{}"
	}

	return strings.Join([]string{"OrderParam", string(data)}, " ")
}

type OrderParamOrder struct {
	value string
}

type OrderParamOrderEnum struct {
	ASC  OrderParamOrder
	DESC OrderParamOrder
}

func GetOrderParamOrderEnum() OrderParamOrderEnum {
	return OrderParamOrderEnum{
		ASC: OrderParamOrder{
			value: "ASC",
		},
		DESC: OrderParamOrder{
			value: "DESC",
		},
	}
}

func (c OrderParamOrder) Value() string {
	return c.value
}

func (c OrderParamOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrderParamOrder) UnmarshalJSON(b []byte) error {
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
