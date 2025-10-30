package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LeakageListInfoAction 规则命中后操作对象
type LeakageListInfoAction struct {

	// 操作类型。   - “block”：过滤。   - “log”：仅记录
	Category LeakageListInfoActionCategory `json:"category"`
}

func (o LeakageListInfoAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeakageListInfoAction struct{}"
	}

	return strings.Join([]string{"LeakageListInfoAction", string(data)}, " ")
}

type LeakageListInfoActionCategory struct {
	value string
}

type LeakageListInfoActionCategoryEnum struct {
	BLOCK LeakageListInfoActionCategory
	LOG   LeakageListInfoActionCategory
}

func GetLeakageListInfoActionCategoryEnum() LeakageListInfoActionCategoryEnum {
	return LeakageListInfoActionCategoryEnum{
		BLOCK: LeakageListInfoActionCategory{
			value: "block",
		},
		LOG: LeakageListInfoActionCategory{
			value: "log",
		},
	}
}

func (c LeakageListInfoActionCategory) Value() string {
	return c.value
}

func (c LeakageListInfoActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LeakageListInfoActionCategory) UnmarshalJSON(b []byte) error {
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
