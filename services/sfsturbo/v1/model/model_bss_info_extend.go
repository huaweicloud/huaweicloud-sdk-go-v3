package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BssInfoExtend 扩容计费参数详情
type BssInfoExtend struct {

	// 是否自动付款
	IsAutoPay *BssInfoExtendIsAutoPay `json:"is_auto_pay,omitempty"`
}

func (o BssInfoExtend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssInfoExtend struct{}"
	}

	return strings.Join([]string{"BssInfoExtend", string(data)}, " ")
}

type BssInfoExtendIsAutoPay struct {
	value int64
}

type BssInfoExtendIsAutoPayEnum struct {
	E_0 BssInfoExtendIsAutoPay
	E_1 BssInfoExtendIsAutoPay
}

func GetBssInfoExtendIsAutoPayEnum() BssInfoExtendIsAutoPayEnum {
	return BssInfoExtendIsAutoPayEnum{
		E_0: BssInfoExtendIsAutoPay{
			value: 0,
		}, E_1: BssInfoExtendIsAutoPay{
			value: 1,
		},
	}
}

func (c BssInfoExtendIsAutoPay) Value() int64 {
	return c.value
}

func (c BssInfoExtendIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssInfoExtendIsAutoPay) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int64")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int64); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int64 error")
	}
}
