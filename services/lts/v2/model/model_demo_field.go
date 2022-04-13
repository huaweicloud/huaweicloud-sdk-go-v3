package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 示例字段对象
type DemoField struct {
	// 字段名称

	FieldName string `json:"field_name"`
	// 字段示例内容

	Content *string `json:"content,omitempty"`
	// 字段数据类型。 可选范围：string、long、float

	Type DemoFieldType `json:"type"`
	// 是否开启快速分析

	IsAnalysis *bool `json:"is_analysis,omitempty"`
	// 手动正则及分隔符方式中字段序号

	Index *int32 `json:"index,omitempty"`
	// 描叙多层级json中字段间的层级关系

	Relation *string `json:"relation,omitempty"`
	// json及nginx方式中字段自定义别名

	UserDefinedName *string `json:"user_defined_name,omitempty"`
}

func (o DemoField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemoField struct{}"
	}

	return strings.Join([]string{"DemoField", string(data)}, " ")
}

type DemoFieldType struct {
	value string
}

type DemoFieldTypeEnum struct {
	STRING DemoFieldType
	LONG   DemoFieldType
	FLOAT  DemoFieldType
}

func GetDemoFieldTypeEnum() DemoFieldTypeEnum {
	return DemoFieldTypeEnum{
		STRING: DemoFieldType{
			value: "string",
		},
		LONG: DemoFieldType{
			value: "long",
		},
		FLOAT: DemoFieldType{
			value: "float",
		},
	}
}

func (c DemoFieldType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DemoFieldType) UnmarshalJSON(b []byte) error {
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
