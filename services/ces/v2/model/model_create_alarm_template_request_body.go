package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAlarmTemplateRequestBody struct {

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]
	TemplateName string `json:"template_name"`

	// 告警模板的描述，长度范围[0,256]，该字段默认值为空字符串
	TemplateDescription *string `json:"template_description,omitempty"`

	// 自定义告警模板类型 0：指标 2： 事件
	TemplateType *CreateAlarmTemplateRequestBodyTemplateType `json:"template_type,omitempty"`

	// 告警模板策略列表
	Policies []Policies `json:"policies"`
}

func (o CreateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateRequestBody", string(data)}, " ")
}

type CreateAlarmTemplateRequestBodyTemplateType struct {
	value int32
}

type CreateAlarmTemplateRequestBodyTemplateTypeEnum struct {
	E_0 CreateAlarmTemplateRequestBodyTemplateType
	E_2 CreateAlarmTemplateRequestBodyTemplateType
}

func GetCreateAlarmTemplateRequestBodyTemplateTypeEnum() CreateAlarmTemplateRequestBodyTemplateTypeEnum {
	return CreateAlarmTemplateRequestBodyTemplateTypeEnum{
		E_0: CreateAlarmTemplateRequestBodyTemplateType{
			value: 0,
		}, E_2: CreateAlarmTemplateRequestBodyTemplateType{
			value: 2,
		},
	}
}

func (c CreateAlarmTemplateRequestBodyTemplateType) Value() int32 {
	return c.value
}

func (c CreateAlarmTemplateRequestBodyTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlarmTemplateRequestBodyTemplateType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
