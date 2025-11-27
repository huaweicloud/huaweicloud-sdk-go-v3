package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecuteDocumentRequestBody struct {

	// 作业版本号，若不传则默认为最新版本
	Version *string `json:"version,omitempty"`

	// 全局参数
	Parameters *[]ExecuteDocumentRequestBodyParameters `json:"parameters,omitempty"`

	// 系统标签列表
	SysTags *[]ExecuteDocumentRequestBodySysTags `json:"sys_tags,omitempty"`

	// 速率控制模式下，批量执行对象的参数名
	TargetParameterName *string `json:"target_parameter_name,omitempty"`

	// 与target_parameter_name搭配使用。选择实例化target_parameter_name参数的方式。
	Targets *[]ExecuteDocumentRequestBodyTargets `json:"targets,omitempty"`

	// 执行的作业类型
	DocumentType *ExecuteDocumentRequestBodyDocumentType `json:"document_type,omitempty"`

	// 执行描述
	Description *string `json:"description,omitempty"`
}

func (o ExecuteDocumentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDocumentRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteDocumentRequestBody", string(data)}, " ")
}

type ExecuteDocumentRequestBodyDocumentType struct {
	value string
}

type ExecuteDocumentRequestBodyDocumentTypeEnum struct {
	PRIVATE ExecuteDocumentRequestBodyDocumentType
	PUBLIC  ExecuteDocumentRequestBodyDocumentType
}

func GetExecuteDocumentRequestBodyDocumentTypeEnum() ExecuteDocumentRequestBodyDocumentTypeEnum {
	return ExecuteDocumentRequestBodyDocumentTypeEnum{
		PRIVATE: ExecuteDocumentRequestBodyDocumentType{
			value: "PRIVATE",
		},
		PUBLIC: ExecuteDocumentRequestBodyDocumentType{
			value: "PUBLIC",
		},
	}
}

func (c ExecuteDocumentRequestBodyDocumentType) Value() string {
	return c.value
}

func (c ExecuteDocumentRequestBodyDocumentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteDocumentRequestBodyDocumentType) UnmarshalJSON(b []byte) error {
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
