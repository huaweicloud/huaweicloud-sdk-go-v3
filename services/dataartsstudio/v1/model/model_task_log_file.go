package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TaskLogFile struct {

	// 文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 文件类型: - DIRECTORY：目录 - FILE：文件
	FileType *TaskLogFileFileType `json:"file_type,omitempty"`

	// 文件大小，单位字节。
	FileSize *int64 `json:"file_size,omitempty"`

	// 文件显示名称。
	DisplayName *string `json:"display_name,omitempty"`
}

func (o TaskLogFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskLogFile struct{}"
	}

	return strings.Join([]string{"TaskLogFile", string(data)}, " ")
}

type TaskLogFileFileType struct {
	value string
}

type TaskLogFileFileTypeEnum struct {
	DIRECTORY TaskLogFileFileType
	FILE      TaskLogFileFileType
}

func GetTaskLogFileFileTypeEnum() TaskLogFileFileTypeEnum {
	return TaskLogFileFileTypeEnum{
		DIRECTORY: TaskLogFileFileType{
			value: "DIRECTORY",
		},
		FILE: TaskLogFileFileType{
			value: "FILE",
		},
	}
}

func (c TaskLogFileFileType) Value() string {
	return c.value
}

func (c TaskLogFileFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskLogFileFileType) UnmarshalJSON(b []byte) error {
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
