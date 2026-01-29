package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSubscriptionOrderResponse Response Object
type CreateSubscriptionOrderResponse struct {

	// 创建或变更订单ID，只有scene为PREPAID时返回有此数据
	OrderId *string `json:"order_id,omitempty"`

	// 订单更新状态，1：变更订单成功，5，订单变更失败
	OrderStatus    *CreateSubscriptionOrderResponseOrderStatus `json:"order_status,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o CreateSubscriptionOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionOrderResponse struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionOrderResponse", string(data)}, " ")
}

type CreateSubscriptionOrderResponseOrderStatus struct {
	value int32
}

type CreateSubscriptionOrderResponseOrderStatusEnum struct {
	E_1 CreateSubscriptionOrderResponseOrderStatus
	E_5 CreateSubscriptionOrderResponseOrderStatus
}

func GetCreateSubscriptionOrderResponseOrderStatusEnum() CreateSubscriptionOrderResponseOrderStatusEnum {
	return CreateSubscriptionOrderResponseOrderStatusEnum{
		E_1: CreateSubscriptionOrderResponseOrderStatus{
			value: 1,
		}, E_5: CreateSubscriptionOrderResponseOrderStatus{
			value: 5,
		},
	}
}

func (c CreateSubscriptionOrderResponseOrderStatus) Value() int32 {
	return c.value
}

func (c CreateSubscriptionOrderResponseOrderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSubscriptionOrderResponseOrderStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
