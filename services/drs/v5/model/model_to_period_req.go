package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ToPeriodReq 任务按需转包周期请求体
type ToPeriodReq struct {

	// 订单类型，取值： - 2：包月 - 3：包年
	PeriodType ToPeriodReqPeriodType `json:"period_type"`

	// 订单周期数
	PeriodNum int32 `json:"period_num"`

	// 是否自动续费
	IsAutoRenew bool `json:"is_auto_renew"`

	// 是否自动支付
	IsAutoPay bool `json:"is_auto_pay"`
}

func (o ToPeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ToPeriodReq struct{}"
	}

	return strings.Join([]string{"ToPeriodReq", string(data)}, " ")
}

type ToPeriodReqPeriodType struct {
	value string
}

type ToPeriodReqPeriodTypeEnum struct {
	E_2 ToPeriodReqPeriodType
	E_3 ToPeriodReqPeriodType
}

func GetToPeriodReqPeriodTypeEnum() ToPeriodReqPeriodTypeEnum {
	return ToPeriodReqPeriodTypeEnum{
		E_2: ToPeriodReqPeriodType{
			value: "2",
		},
		E_3: ToPeriodReqPeriodType{
			value: "3",
		},
	}
}

func (c ToPeriodReqPeriodType) Value() string {
	return c.value
}

func (c ToPeriodReqPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ToPeriodReqPeriodType) UnmarshalJSON(b []byte) error {
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
