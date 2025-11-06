package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StructConfig 结构化配置参数体
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

func (c StructConfigTemplateType) Value() string {
	return c.value
}

func (c StructConfigTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StructConfigTemplateType) UnmarshalJSON(b []byte) error {
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
