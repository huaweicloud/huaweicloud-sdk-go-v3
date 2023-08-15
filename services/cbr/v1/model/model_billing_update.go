package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BillingUpdate struct {

	// 存储库规格
	ConsistentLevel *BillingUpdateConsistentLevel `json:"consistent_level,omitempty"`

	// 存储库大小，单位为GB
	Size *int32 `json:"size,omitempty"`
}

func (o BillingUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingUpdate struct{}"
	}

	return strings.Join([]string{"BillingUpdate", string(data)}, " ")
}

type BillingUpdateConsistentLevel struct {
	value string
}

type BillingUpdateConsistentLevelEnum struct {
	APP_CONSISTENT   BillingUpdateConsistentLevel
	CRASH_CONSISTENT BillingUpdateConsistentLevel
}

func GetBillingUpdateConsistentLevelEnum() BillingUpdateConsistentLevelEnum {
	return BillingUpdateConsistentLevelEnum{
		APP_CONSISTENT: BillingUpdateConsistentLevel{
			value: "app_consistent",
		},
		CRASH_CONSISTENT: BillingUpdateConsistentLevel{
			value: "crash_consistent",
		},
	}
}

func (c BillingUpdateConsistentLevel) Value() string {
	return c.value
}

func (c BillingUpdateConsistentLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingUpdateConsistentLevel) UnmarshalJSON(b []byte) error {
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
