package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 结构化模板实体
type StructTemplate struct {
	// 项目id

	ProjectId string `json:"project_id"`
	// 模板名称

	TemplateName string `json:"template_name"`
	// 模板类型，regex,json,split,nginx

	TemplateType StructTemplateTemplateType `json:"template_type"`
	// 示例日志

	DemoLog string `json:"demo_log"`
	// 示例字段数组

	DemoFields []DemoField `json:"demo_fields"`
	// Tag字段数组

	TagFields []TagFieldNew `json:"tag_fields"`

	Rule *TemplateRule `json:"rule"`
	// 示例日志标签

	DemoLabel *string `json:"demo_label,omitempty"`
	// 创建时间

	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o StructTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructTemplate struct{}"
	}

	return strings.Join([]string{"StructTemplate", string(data)}, " ")
}

type StructTemplateTemplateType struct {
	value string
}

type StructTemplateTemplateTypeEnum struct {
	REGEX StructTemplateTemplateType
	JSON  StructTemplateTemplateType
	SPLIT StructTemplateTemplateType
	NGINX StructTemplateTemplateType
}

func GetStructTemplateTemplateTypeEnum() StructTemplateTemplateTypeEnum {
	return StructTemplateTemplateTypeEnum{
		REGEX: StructTemplateTemplateType{
			value: "regex",
		},
		JSON: StructTemplateTemplateType{
			value: "json",
		},
		SPLIT: StructTemplateTemplateType{
			value: "split",
		},
		NGINX: StructTemplateTemplateType{
			value: "nginx",
		},
	}
}

func (c StructTemplateTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StructTemplateTemplateType) UnmarshalJSON(b []byte) error {
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
