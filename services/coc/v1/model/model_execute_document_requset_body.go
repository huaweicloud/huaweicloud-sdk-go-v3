package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecuteDocumentRequsetBody struct {

	// 作业版本号，若不传则默认为最新版本
	Version *string `json:"version,omitempty"`

	// 全局参数
	Parameters *[]ExecuteDocumentRequsetBodyParameters `json:"parameters,omitempty"`

	// 系统标签列表
	SysTags *[]ExecuteDocumentRequsetBodySysTags `json:"sys_tags,omitempty"`

	// 速率控制模式下，批量执行对象的参数名
	TargetParameterName *string `json:"target_parameter_name,omitempty"`

	// 与target_parameter_name搭配使用。选择实例化target_parameter_name参数的方式。
	Targets *[]ExecuteDocumentRequsetBodyTargets `json:"targets,omitempty"`

	// 执行的作业类型
	DocumentType *ExecuteDocumentRequsetBodyDocumentType `json:"document_type,omitempty"`

	// 执行描述
	Description *string `json:"description,omitempty"`
}

func (o ExecuteDocumentRequsetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequsetBody struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequsetBody", string(data)}, " ")
}

type ExecuteDocumentRequsetBodyDocumentType struct {
	value string
}

type ExecuteDocumentRequsetBodyDocumentTypeEnum struct {
	PRIVATE ExecuteDocumentRequsetBodyDocumentType
	PUBLIC  ExecuteDocumentRequsetBodyDocumentType
}

func GetExecuteDocumentRequsetBodyDocumentTypeEnum() ExecuteDocumentRequsetBodyDocumentTypeEnum {
	return ExecuteDocumentRequsetBodyDocumentTypeEnum{
		PRIVATE: ExecuteDocumentRequsetBodyDocumentType{
			value: "PRIVATE",
		},
		PUBLIC: ExecuteDocumentRequsetBodyDocumentType{
			value: "PUBLIC",
		},
	}
}

func (c ExecuteDocumentRequsetBodyDocumentType) Value() string {
	return c.value
}

func (c ExecuteDocumentRequsetBodyDocumentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteDocumentRequsetBodyDocumentType) UnmarshalJSON(b []byte) error {
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
