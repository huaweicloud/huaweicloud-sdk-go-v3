package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAlarmTemplateRequestBody struct {

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]
	TemplateName string `json:"template_name"`

	// **参数解释**： 自定义告警模板类型 **约束限制**： 当template_type为0或者不选时，policies中的dimension_name必填。当template_type为2时，dimension_name为空。 **取值范围**： 枚举值。0：指标；2： 事件。 **默认取值**： 0
	TemplateType *UpdateAlarmTemplateRequestBodyTemplateType `json:"template_type,omitempty"`

	// 告警模板的描述，长度范围[0,256]，该字段默认值为空字符串
	TemplateDescription *string `json:"template_description,omitempty"`

	// 告警模板策略列表
	Policies []Policies `json:"policies"`
}

func (o UpdateAlarmTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequestBody", string(data)}, " ")
}

type UpdateAlarmTemplateRequestBodyTemplateType struct {
	value int32
}

type UpdateAlarmTemplateRequestBodyTemplateTypeEnum struct {
	E_0 UpdateAlarmTemplateRequestBodyTemplateType
	E_2 UpdateAlarmTemplateRequestBodyTemplateType
}

func GetUpdateAlarmTemplateRequestBodyTemplateTypeEnum() UpdateAlarmTemplateRequestBodyTemplateTypeEnum {
	return UpdateAlarmTemplateRequestBodyTemplateTypeEnum{
		E_0: UpdateAlarmTemplateRequestBodyTemplateType{
			value: 0,
		}, E_2: UpdateAlarmTemplateRequestBodyTemplateType{
			value: 2,
		},
	}
}

func (c UpdateAlarmTemplateRequestBodyTemplateType) Value() int32 {
	return c.value
}

func (c UpdateAlarmTemplateRequestBodyTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlarmTemplateRequestBodyTemplateType) UnmarshalJSON(b []byte) error {
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
