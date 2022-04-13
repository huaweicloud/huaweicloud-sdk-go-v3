package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 结构化配置参数体
type StructConfig struct {
	// 日志组ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。

	LogGroupId string `json:"log_group_id"`
	// 日志流ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。

	LogStreamId string `json:"log_stream_id"`
	// 所用模板id。当使用系统模板时，当前属性可以为空

	TemplateId string `json:"template_id"`
	// 所用模板名称，会对模板名称及id进行校验

	TemplateName string `json:"template_name"`
	// 所用模板类型，分为built_in及custom两种类型，对应系统模板和自定义模板，系统模板分为CTS，VPC和ELB三种。

	TemplateType StructConfigTemplateType `json:"template_type"`
	// 示例字段数组，只需要填写与模板中is_analysis状态不同的字段

	DemoFields *[]FieldModel `json:"demo_fields,omitempty"`
	// Tag字段数组，只需要填写与模板中is_analysis状态不同的字段

	TagFields *[]FieldModel `json:"tag_fields,omitempty"`
	// 是否开启demo_fields和tag_fields快速分析,为true时，所有的demo_fields和tag_fields全部字段均打开快速分析;不填或者为false，以模板中的demo_fields和tag_fields中的is_analysis决定是否开启快速分析。

	QuickAnalysis *bool `json:"quick_analysis,omitempty"`
}

func (o StructConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructConfig struct{}"
	}

	return strings.Join([]string{"StructConfig", string(data)}, " ")
}

type StructConfigTemplateType struct {
	value string
}

type StructConfigTemplateTypeEnum struct {
	BUILT_IN StructConfigTemplateType
	CUSTOM   StructConfigTemplateType
}

func GetStructConfigTemplateTypeEnum() StructConfigTemplateTypeEnum {
	return StructConfigTemplateTypeEnum{
		BUILT_IN: StructConfigTemplateType{
			value: "built_in",
		},
		CUSTOM: StructConfigTemplateType{
			value: "custom",
		},
	}
}

func (c StructConfigTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StructConfigTemplateType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
