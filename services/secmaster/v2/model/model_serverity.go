package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Serverity **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
type Serverity struct {
	value string
}

type ServerityEnum struct {
	TIPS   Serverity
	LOW    Serverity
	MEDIUM Serverity
	HIGH   Serverity
	FATAL  Serverity
}

func GetServerityEnum() ServerityEnum {
	return ServerityEnum{
		TIPS: Serverity{
			value: "TIPS",
		},
		LOW: Serverity{
			value: "LOW",
		},
		MEDIUM: Serverity{
			value: "MEDIUM",
		},
		HIGH: Serverity{
			value: "HIGH",
		},
		FATAL: Serverity{
			value: "FATAL",
		},
	}
}

func (c Serverity) Value() string {
	return c.value
}

func (c Serverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Serverity) UnmarshalJSON(b []byte) error {
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
