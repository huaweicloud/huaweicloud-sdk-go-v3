package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModelFileSource 模型数据文件来源
type ModelFileSource struct {
	value string
}

type ModelFileSourceEnum struct {
	PUBLIC  ModelFileSource
	PRIVATE ModelFileSource
}

func GetModelFileSourceEnum() ModelFileSourceEnum {
	return ModelFileSourceEnum{
		PUBLIC: ModelFileSource{
			value: "public",
		},
		PRIVATE: ModelFileSource{
			value: "private",
		},
	}
}

func (c ModelFileSource) Value() string {
	return c.value
}

func (c ModelFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModelFileSource) UnmarshalJSON(b []byte) error {
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
