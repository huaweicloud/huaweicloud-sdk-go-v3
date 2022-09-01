package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 订阅批量操作请求
type SubscriptionOperateReq struct {

	// 订阅对象ID列表，单次批量操作最多支持10个订阅
	SubscriptionIds *[]string `json:"subscription_ids,omitempty" xml:"subscription_ids"`

	// 操作类型
	Operation *SubscriptionOperateReqOperation `json:"operation,omitempty" xml:"operation"`
}

func (o SubscriptionOperateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionOperateReq struct{}"
	}

	return strings.Join([]string{"SubscriptionOperateReq", string(data)}, " ")
}

type SubscriptionOperateReqOperation struct {
	value string
}

type SubscriptionOperateReqOperationEnum struct {
	DISABLE SubscriptionOperateReqOperation
	ENABLE  SubscriptionOperateReqOperation
}

func GetSubscriptionOperateReqOperationEnum() SubscriptionOperateReqOperationEnum {
	return SubscriptionOperateReqOperationEnum{
		DISABLE: SubscriptionOperateReqOperation{
			value: "DISABLE",
		},
		ENABLE: SubscriptionOperateReqOperation{
			value: "ENABLE",
		},
	}
}

func (c SubscriptionOperateReqOperation) Value() string {
	return c.value
}

func (c SubscriptionOperateReqOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionOperateReqOperation) UnmarshalJSON(b []byte) error {
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
