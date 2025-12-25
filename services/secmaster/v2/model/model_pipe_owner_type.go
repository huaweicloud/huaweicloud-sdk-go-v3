package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipeOwnerType **参数解释**: 管道所有者类型 - SYSTEM 系统 - USER 用户 - CLOUD_LOG 云日志  **约束限制** 不涉及 **取值范围**: - SYSTEM - USER - CLOUD_LOG  **默认值** 不涉及
type PipeOwnerType struct {
	value string
}

type PipeOwnerTypeEnum struct {
	SYSTEM    PipeOwnerType
	USER      PipeOwnerType
	CLOUD_LOG PipeOwnerType
}

func GetPipeOwnerTypeEnum() PipeOwnerTypeEnum {
	return PipeOwnerTypeEnum{
		SYSTEM: PipeOwnerType{
			value: "SYSTEM",
		},
		USER: PipeOwnerType{
			value: "USER",
		},
		CLOUD_LOG: PipeOwnerType{
			value: "CLOUD_LOG",
		},
	}
}

func (c PipeOwnerType) Value() string {
	return c.value
}

func (c PipeOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipeOwnerType) UnmarshalJSON(b []byte) error {
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
