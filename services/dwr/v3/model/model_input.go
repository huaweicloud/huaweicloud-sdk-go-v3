package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Input 创建算子时用户可修改参数列表
type Input struct {

	// 输入参数名称，由小写字母、数字和中划线“-”组成
	ParameterName string `json:"parameter_name"`

	// 参数类型。可为string，integer，float，boolean，list，map。type为list类型时，value_type必填
	Type InputType `json:"type"`

	ValueType *InputParaValueType `json:"value_type,omitempty"`

	// 当多个相同action template在一个工作流时，增加字段做国际化。 由小写字母、数字和中划线“-”组成
	TemplateParameterName *string `json:"template_parameter_name,omitempty"`

	// Input结构体参数类型。支持string,integer,float,boolean,list,map
	ParameterValue *string `json:"parameter_value,omitempty"`

	// 参数项描述信息。
	Description *string `json:"description,omitempty"`

	// 默认值信息可在创建工作流实例时由外部输入替换；若未填写默认值，外部输入将必须填写这个参数的值。 注：默认值的类型和定义的参数类型必须统一。如果出现不一致，解析器可能会进行自动转换而导致出现与预期不符合的情况。
	Default *string `json:"default,omitempty"`

	// 参数的标签，此处定义的标签可在创建堆栈时进行分类展示。
	Label *string `json:"label,omitempty"`

	// 约束条件有以下几种，一个输入参数对每一种条件都只能定义一个规则。约束的多个条件中只要有一条不满足，即将认定参数非法。 equal：约定参数的value值必须等于特定值。 valid_values：参数的有效值，定义一个数组。 regex：参数需要满足某个正则条件，必须是字符串类型才可以进行匹配。 invalid_values：参数的无效值范围，如果参数值定义在其中，将会认为无效而报错。
	Constraints *interface{} `json:"constraints,omitempty"`

	// 输入参数的invisible设置为true时，返回值为******。
	Invisible *bool `json:"invisible,omitempty"`
}

func (o Input) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Input struct{}"
	}

	return strings.Join([]string{"Input", string(data)}, " ")
}

type InputType struct {
	value string
}

type InputTypeEnum struct {
	STRING  InputType
	INTEGER InputType
	FLOAT   InputType
	BOOLEAN InputType
	LIST    InputType
	MAP     InputType
}

func GetInputTypeEnum() InputTypeEnum {
	return InputTypeEnum{
		STRING: InputType{
			value: "string",
		},
		INTEGER: InputType{
			value: "integer",
		},
		FLOAT: InputType{
			value: "float",
		},
		BOOLEAN: InputType{
			value: "boolean",
		},
		LIST: InputType{
			value: "list",
		},
		MAP: InputType{
			value: "map",
		},
	}
}

func (c InputType) Value() string {
	return c.value
}

func (c InputType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InputType) UnmarshalJSON(b []byte) error {
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
