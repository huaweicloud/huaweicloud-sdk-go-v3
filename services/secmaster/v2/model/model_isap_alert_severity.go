package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapAlertSeverity **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
type IsapAlertSeverity struct {
	value string
}

type IsapAlertSeverityEnum struct {
	TIPS   IsapAlertSeverity
	LOW    IsapAlertSeverity
	MEDIUM IsapAlertSeverity
	HIGH   IsapAlertSeverity
	FATAL  IsapAlertSeverity
}

func GetIsapAlertSeverityEnum() IsapAlertSeverityEnum {
	return IsapAlertSeverityEnum{
		TIPS: IsapAlertSeverity{
			value: "TIPS",
		},
		LOW: IsapAlertSeverity{
			value: "LOW",
		},
		MEDIUM: IsapAlertSeverity{
			value: "MEDIUM",
		},
		HIGH: IsapAlertSeverity{
			value: "HIGH",
		},
		FATAL: IsapAlertSeverity{
			value: "FATAL",
		},
	}
}

func (c IsapAlertSeverity) Value() string {
	return c.value
}

func (c IsapAlertSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapAlertSeverity) UnmarshalJSON(b []byte) error {
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
