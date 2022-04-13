package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 更新或者查询结构化模板对象
type StructTemplateModel struct {
	// 项目id

	ProjectId string `json:"project_id"`
	// 模板名称

	TemplateName string `json:"template_name"`
	// 模板类型，regex,json,split,nginx

	TemplateType StructTemplateModelTemplateType `json:"template_type"`
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
	// 模板id

	Id *string `json:"id,omitempty"`
}

func (o StructTemplateModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructTemplateModel struct{}"
	}

	return strings.Join([]string{"StructTemplateModel", string(data)}, " ")
}

type StructTemplateModelTemplateType struct {
	value string
}

type StructTemplateModelTemplateTypeEnum struct {
	REGEX StructTemplateModelTemplateType
	JSON  StructTemplateModelTemplateType
	SPLIT StructTemplateModelTemplateType
	NGINX StructTemplateModelTemplateType
}

func GetStructTemplateModelTemplateTypeEnum() StructTemplateModelTemplateTypeEnum {
	return StructTemplateModelTemplateTypeEnum{
		REGEX: StructTemplateModelTemplateType{
			value: "regex",
		},
		JSON: StructTemplateModelTemplateType{
			value: "json",
		},
		SPLIT: StructTemplateModelTemplateType{
			value: "split",
		},
		NGINX: StructTemplateModelTemplateType{
			value: "nginx",
		},
	}
}

func (c StructTemplateModelTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StructTemplateModelTemplateType) UnmarshalJSON(b []byte) error {
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
