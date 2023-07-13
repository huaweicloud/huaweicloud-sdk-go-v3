package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FunctionResourceUri 函数包信息
type FunctionResourceUri struct {

	// 函数包类型
	Type FunctionResourceUriType `json:"type"`

	// 函数包地址信息
	Uri string `json:"uri"`
}

func (o FunctionResourceUri) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionResourceUri struct{}"
	}

	return strings.Join([]string{"FunctionResourceUri", string(data)}, " ")
}

type FunctionResourceUriType struct {
	value string
}

type FunctionResourceUriTypeEnum struct {
	JAR     FunctionResourceUriType
	FILE    FunctionResourceUriType
	ARCHIVE FunctionResourceUriType
}

func GetFunctionResourceUriTypeEnum() FunctionResourceUriTypeEnum {
	return FunctionResourceUriTypeEnum{
		JAR: FunctionResourceUriType{
			value: "JAR",
		},
		FILE: FunctionResourceUriType{
			value: "FILE",
		},
		ARCHIVE: FunctionResourceUriType{
			value: "ARCHIVE",
		},
	}
}

func (c FunctionResourceUriType) Value() string {
	return c.value
}

func (c FunctionResourceUriType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FunctionResourceUriType) UnmarshalJSON(b []byte) error {
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
