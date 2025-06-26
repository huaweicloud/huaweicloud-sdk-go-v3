package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DiagnosisTaskSubmitBody 诊断任务提交请求的结构体
type DiagnosisTaskSubmitBody struct {

	// 待诊断的实例ID
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 被诊断实例的类型
	Type *DiagnosisTaskSubmitBodyType `json:"type,omitempty"`

	// 适用于RDS、DMS、DCS、ELB等的额外参数（JSON字符串），该数组长度应与提交的实例数组长度对应
	ExtraProperties *[]string `json:"extra_properties,omitempty"`
}

func (o DiagnosisTaskSubmitBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisTaskSubmitBody struct{}"
	}

	return strings.Join([]string{"DiagnosisTaskSubmitBody", string(data)}, " ")
}

type DiagnosisTaskSubmitBodyType struct {
	value string
}

type DiagnosisTaskSubmitBodyTypeEnum struct {
	ECS DiagnosisTaskSubmitBodyType
	RDS DiagnosisTaskSubmitBodyType
	DCS DiagnosisTaskSubmitBodyType
	DMS DiagnosisTaskSubmitBodyType
	ELB DiagnosisTaskSubmitBodyType
}

func GetDiagnosisTaskSubmitBodyTypeEnum() DiagnosisTaskSubmitBodyTypeEnum {
	return DiagnosisTaskSubmitBodyTypeEnum{
		ECS: DiagnosisTaskSubmitBodyType{
			value: "ECS",
		},
		RDS: DiagnosisTaskSubmitBodyType{
			value: "RDS",
		},
		DCS: DiagnosisTaskSubmitBodyType{
			value: "DCS",
		},
		DMS: DiagnosisTaskSubmitBodyType{
			value: "DMS",
		},
		ELB: DiagnosisTaskSubmitBodyType{
			value: "ELB",
		},
	}
}

func (c DiagnosisTaskSubmitBodyType) Value() string {
	return c.value
}

func (c DiagnosisTaskSubmitBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiagnosisTaskSubmitBodyType) UnmarshalJSON(b []byte) error {
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
