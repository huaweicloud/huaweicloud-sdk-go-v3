package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type WafProductInfo struct {

	// waf规格   -  professional：标准   - enterprise：专业   - ultimate：企业版
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// **参数解释：** 订购周期类型标识，用于指定套餐的订购时间周期单位 **约束限制：** 不涉及 **取值范围：**  - month:月  - year:年 **默认取值：** 不涉及
	PeriodType *WafProductInfoPeriodType `json:"period_type,omitempty"`

	// 订购周期数
	PeriodNum *int32 `json:"period_num,omitempty"`
}

func (o WafProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafProductInfo struct{}"
	}

	return strings.Join([]string{"WafProductInfo", string(data)}, " ")
}

type WafProductInfoPeriodType struct {
	value string
}

type WafProductInfoPeriodTypeEnum struct {
	MONTH WafProductInfoPeriodType
	YEAR  WafProductInfoPeriodType
}

func GetWafProductInfoPeriodTypeEnum() WafProductInfoPeriodTypeEnum {
	return WafProductInfoPeriodTypeEnum{
		MONTH: WafProductInfoPeriodType{
			value: "month",
		},
		YEAR: WafProductInfoPeriodType{
			value: "year",
		},
	}
}

func (c WafProductInfoPeriodType) Value() string {
	return c.value
}

func (c WafProductInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WafProductInfoPeriodType) UnmarshalJSON(b []byte) error {
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
