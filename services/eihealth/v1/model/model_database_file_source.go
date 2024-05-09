package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DatabaseFileSource 数据库文件来源
type DatabaseFileSource struct {
	value string
}

type DatabaseFileSourceEnum struct {
	PUBLIC  DatabaseFileSource
	PRIVATE DatabaseFileSource
}

func GetDatabaseFileSourceEnum() DatabaseFileSourceEnum {
	return DatabaseFileSourceEnum{
		PUBLIC: DatabaseFileSource{
			value: "public",
		},
		PRIVATE: DatabaseFileSource{
			value: "private",
		},
	}
}

func (c DatabaseFileSource) Value() string {
	return c.value
}

func (c DatabaseFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseFileSource) UnmarshalJSON(b []byte) error {
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
