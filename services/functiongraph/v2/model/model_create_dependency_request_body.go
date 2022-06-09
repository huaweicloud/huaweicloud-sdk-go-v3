package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateDependencyRequestBody struct {

	// depend_type为zip类型时必填，为文件流格式,需要base64编码zip文件。
	DependFile *string `json:"depend_file,omitempty"`

	// depend_type为obs类型时，依赖包在obs的存储地址。
	DependLink *string `json:"depend_link,omitempty"`

	// 导入类型,目前支持obs和zip。
	DependType string `json:"depend_type"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本
	Runtime CreateDependencyRequestBodyRuntime `json:"runtime"`

	// 依赖包名称。必须以大、小写字母开头，以字母或数字结尾，只能由字母、数字、下划线、点和中划线组成，长度不超过96个字符。
	Name string `json:"name"`

	// 依赖包描述，不超过512个字符。
	Description *string `json:"description,omitempty"`
}

func (o CreateDependencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDependencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDependencyRequestBody", string(data)}, " ")
}

type CreateDependencyRequestBodyRuntime struct {
	value string
}

type CreateDependencyRequestBodyRuntimeEnum struct {
	JAVA8           CreateDependencyRequestBodyRuntime
	JAVA11          CreateDependencyRequestBodyRuntime
	NODE_JS6_10     CreateDependencyRequestBodyRuntime
	NODE_JS8_10     CreateDependencyRequestBodyRuntime
	NODE_JS10_16    CreateDependencyRequestBodyRuntime
	NODE_JS12_13    CreateDependencyRequestBodyRuntime
	NODE_JS14_18    CreateDependencyRequestBodyRuntime
	PYTHON2_7       CreateDependencyRequestBodyRuntime
	PYTHON3_6       CreateDependencyRequestBodyRuntime
	GO1_8           CreateDependencyRequestBodyRuntime
	GO1_X           CreateDependencyRequestBodyRuntime
	C__NET_CORE_2_0 CreateDependencyRequestBodyRuntime
	C__NET_CORE_2_1 CreateDependencyRequestBodyRuntime
	C__NET_CORE_3_1 CreateDependencyRequestBodyRuntime
	PHP7_3          CreateDependencyRequestBodyRuntime
	PYTHON3_9       CreateDependencyRequestBodyRuntime
}

func GetCreateDependencyRequestBodyRuntimeEnum() CreateDependencyRequestBodyRuntimeEnum {
	return CreateDependencyRequestBodyRuntimeEnum{
		JAVA8: CreateDependencyRequestBodyRuntime{
			value: "Java8",
		},
		JAVA11: CreateDependencyRequestBodyRuntime{
			value: "Java11",
		},
		NODE_JS6_10: CreateDependencyRequestBodyRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: CreateDependencyRequestBodyRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: CreateDependencyRequestBodyRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: CreateDependencyRequestBodyRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: CreateDependencyRequestBodyRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: CreateDependencyRequestBodyRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: CreateDependencyRequestBodyRuntime{
			value: "Python3.6",
		},
		GO1_8: CreateDependencyRequestBodyRuntime{
			value: "Go1.8",
		},
		GO1_X: CreateDependencyRequestBodyRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: CreateDependencyRequestBodyRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: CreateDependencyRequestBodyRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: CreateDependencyRequestBodyRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: CreateDependencyRequestBodyRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: CreateDependencyRequestBodyRuntime{
			value: "Python3.9",
		},
	}
}

func (c CreateDependencyRequestBodyRuntime) Value() string {
	return c.value
}

func (c CreateDependencyRequestBodyRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDependencyRequestBodyRuntime) UnmarshalJSON(b []byte) error {
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
