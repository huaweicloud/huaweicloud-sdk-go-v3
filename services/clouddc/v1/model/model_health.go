package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Health 硬件健康状态： OK：健康 Warning：警告 Critical：严重 Unknown：未知
type Health struct {
	value string
}

type HealthEnum struct {
	OK       Health
	WARNING  Health
	CRITICAL Health
	UNKNOWN  Health
}

func GetHealthEnum() HealthEnum {
	return HealthEnum{
		OK: Health{
			value: "OK",
		},
		WARNING: Health{
			value: "Warning",
		},
		CRITICAL: Health{
			value: "Critical",
		},
		UNKNOWN: Health{
			value: "Unknown",
		},
	}
}

func (c Health) Value() string {
	return c.value
}

func (c Health) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Health) UnmarshalJSON(b []byte) error {
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
