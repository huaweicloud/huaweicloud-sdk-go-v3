package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSubscriptionOrderResponse Response Object
type UpdateSubscriptionOrderResponse struct {

	// 创建或变更订单ID，只有scene为PREPAID时返回有此数据
	OrderId *string `json:"order_id,omitempty"`

	// 订单更新状态，1：变更订单成功，5，订单变更失败
	OrderStatus    *UpdateSubscriptionOrderResponseOrderStatus `json:"order_status,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o UpdateSubscriptionOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionOrderResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionOrderResponse", string(data)}, " ")
}

type UpdateSubscriptionOrderResponseOrderStatus struct {
	value int32
}

type UpdateSubscriptionOrderResponseOrderStatusEnum struct {
	E_1 UpdateSubscriptionOrderResponseOrderStatus
	E_5 UpdateSubscriptionOrderResponseOrderStatus
}

func GetUpdateSubscriptionOrderResponseOrderStatusEnum() UpdateSubscriptionOrderResponseOrderStatusEnum {
	return UpdateSubscriptionOrderResponseOrderStatusEnum{
		E_1: UpdateSubscriptionOrderResponseOrderStatus{
			value: 1,
		}, E_5: UpdateSubscriptionOrderResponseOrderStatus{
			value: 5,
		},
	}
}

func (c UpdateSubscriptionOrderResponseOrderStatus) Value() int32 {
	return c.value
}

func (c UpdateSubscriptionOrderResponseOrderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubscriptionOrderResponseOrderStatus) UnmarshalJSON(b []byte) error {
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
