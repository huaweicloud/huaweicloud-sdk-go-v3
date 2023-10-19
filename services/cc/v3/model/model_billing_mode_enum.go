package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BillingModeEnum 带宽包实例在大陆站或国际站的计费方式。 1：大陆站包周期 2：国际站包周期 3：大陆站按需计费 4：国际站按需计费 5：大陆站按95方式计费 6：国际站按95方式计费
type BillingModeEnum struct {
	value int32
}

type BillingModeEnumEnum struct {
	E_1 BillingModeEnum
	E_2 BillingModeEnum
	E_3 BillingModeEnum
	E_4 BillingModeEnum
	E_5 BillingModeEnum
	E_6 BillingModeEnum
}

func GetBillingModeEnumEnum() BillingModeEnumEnum {
	return BillingModeEnumEnum{
		E_1: BillingModeEnum{
			value: 1,
		}, E_2: BillingModeEnum{
			value: 2,
		}, E_3: BillingModeEnum{
			value: 3,
		}, E_4: BillingModeEnum{
			value: 4,
		}, E_5: BillingModeEnum{
			value: 5,
		}, E_6: BillingModeEnum{
			value: 6,
		},
	}
}

func (c BillingModeEnum) Value() int32 {
	return c.value
}

func (c BillingModeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingModeEnum) UnmarshalJSON(b []byte) error {
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
