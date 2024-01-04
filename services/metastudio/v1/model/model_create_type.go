package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateType 任务创建方式。 * PACKAGE: 使用一个zip包包含所有数据 * SEGMENT: 逐句上传数据
type CreateType struct {
	value string
}

type CreateTypeEnum struct {
	PACKAGE CreateType
	SEGMENT CreateType
}

func GetCreateTypeEnum() CreateTypeEnum {
	return CreateTypeEnum{
		PACKAGE: CreateType{
			value: "PACKAGE",
		},
		SEGMENT: CreateType{
			value: "SEGMENT",
		},
	}
}

func (c CreateType) Value() string {
	return c.value
}

func (c CreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateType) UnmarshalJSON(b []byte) error {
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
