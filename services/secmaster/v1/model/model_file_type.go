package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FileType **参数解释**: 文件类型 - JVM JVM配置文件类型 - LOG4J2 Log4j2日志配置文件类型 - YML YAML配置文件类型  **约束限制** 不涉及 **取值范围**: - JVM - LOG4J2 - YML   **默认值** 不涉及
type FileType struct {
	value string
}

type FileTypeEnum struct {
	JVM     FileType
	LOG4_J2 FileType
	YML     FileType
}

func GetFileTypeEnum() FileTypeEnum {
	return FileTypeEnum{
		JVM: FileType{
			value: "JVM",
		},
		LOG4_J2: FileType{
			value: "LOG4J2",
		},
		YML: FileType{
			value: "YML",
		},
	}
}

func (c FileType) Value() string {
	return c.value
}

func (c FileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FileType) UnmarshalJSON(b []byte) error {
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
