package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BssParamEntity struct {

	// 功能说明：下单订购后，是否自动从客户的账户中支付；默认是“不自动支付” 。  取值范围： - true：是（自动支付，从账户余额自动扣费） - false：否（默认值，只提交订单不支付，需要客户手动去支付）  约束： 自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择不自动支付，然后在用户费用中心，选择代金券支付。  **如果没有设置成自动支付，即设置为false时，在创建实例之后，实例状态为“支付中”，用户必须在“费用中心 > 我的订单”，完成订单支付，否则订单一直在支付中，实例没有创建成功**。
	IsAutoPay *BssParamEntityIsAutoPay `json:"is_auto_pay,omitempty"`
}

func (o BssParamEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParamEntity struct{}"
	}

	return strings.Join([]string{"BssParamEntity", string(data)}, " ")
}

type BssParamEntityIsAutoPay struct {
	value string
}

type BssParamEntityIsAutoPayEnum struct {
	TRUE  BssParamEntityIsAutoPay
	FALSE BssParamEntityIsAutoPay
}

func GetBssParamEntityIsAutoPayEnum() BssParamEntityIsAutoPayEnum {
	return BssParamEntityIsAutoPayEnum{
		TRUE: BssParamEntityIsAutoPay{
			value: "true",
		},
		FALSE: BssParamEntityIsAutoPay{
			value: "false",
		},
	}
}

func (c BssParamEntityIsAutoPay) Value() string {
	return c.value
}

func (c BssParamEntityIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamEntityIsAutoPay) UnmarshalJSON(b []byte) error {
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
