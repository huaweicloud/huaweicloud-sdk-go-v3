package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Tag字段对象
type TagFieldNew struct {
	// 字段名称

	FieldName string `json:"field_name"`
	// 字段示例内容

	Content *string `json:"content,omitempty"`
	// 字段数据类型。 可选范围：string、long、float

	Type TagFieldNewType `json:"type"`
	// 是否开启快速分析

	IsAnalysis *bool `json:"is_analysis,omitempty"`
	// 序号，从0开始

	Index *int32 `json:"index,omitempty"`
}

func (o TagFieldNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFieldNew struct{}"
	}

	return strings.Join([]string{"TagFieldNew", string(data)}, " ")
}

type TagFieldNewType struct {
	value string
}

type TagFieldNewTypeEnum struct {
	STRING TagFieldNewType
	LONG   TagFieldNewType
	FLOAT  TagFieldNewType
}

func GetTagFieldNewTypeEnum() TagFieldNewTypeEnum {
	return TagFieldNewTypeEnum{
		STRING: TagFieldNewType{
			value: "string",
		},
		LONG: TagFieldNewType{
			value: "long",
		},
		FLOAT: TagFieldNewType{
			value: "float",
		},
	}
}

func (c TagFieldNewType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TagFieldNewType) UnmarshalJSON(b []byte) error {
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
