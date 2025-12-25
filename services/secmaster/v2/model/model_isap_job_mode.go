package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapJobMode **参数解释**: 作业模式 - STREAMING 流式处理 - BATCH 批处理 - SEARCH 检索  **约束限制** 不涉及 **取值范围**: - STREAMING - BATCH - SEARCH  **默认值** 不涉及
type IsapJobMode struct {
	value string
}

type IsapJobModeEnum struct {
	STREAMING IsapJobMode
	BATCH     IsapJobMode
	SEARCH    IsapJobMode
}

func GetIsapJobModeEnum() IsapJobModeEnum {
	return IsapJobModeEnum{
		STREAMING: IsapJobMode{
			value: "STREAMING",
		},
		BATCH: IsapJobMode{
			value: "BATCH",
		},
		SEARCH: IsapJobMode{
			value: "SEARCH",
		},
	}
}

func (c IsapJobMode) Value() string {
	return c.value
}

func (c IsapJobMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapJobMode) UnmarshalJSON(b []byte) error {
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
