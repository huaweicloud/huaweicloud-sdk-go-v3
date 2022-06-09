package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateDependencyResponse struct {

	// 依赖包ID。
	Id *string `json:"id,omitempty"`

	// 依赖包拥有者。
	Owner *string `json:"owner,omitempty"`

	// 依赖包在obs的存储地址。
	Link *string `json:"link,omitempty"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本
	Runtime *CreateDependencyResponseRuntime `json:"runtime,omitempty"`

	// 依赖包唯一标志。
	Etag *string `json:"etag,omitempty"`

	// 依赖包大小。
	Size *int64 `json:"size,omitempty"`

	// 依赖包名。
	Name *string `json:"name,omitempty"`

	// 依赖包描述。
	Description *string `json:"description,omitempty"`

	// 依赖包文件名。
	FileName       *string `json:"file_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDependencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDependencyResponse struct{}"
	}

	return strings.Join([]string{"CreateDependencyResponse", string(data)}, " ")
}

type CreateDependencyResponseRuntime struct {
	value string
}

type CreateDependencyResponseRuntimeEnum struct {
	JAVA8           CreateDependencyResponseRuntime
	JAVA11          CreateDependencyResponseRuntime
	NODE_JS6_10     CreateDependencyResponseRuntime
	NODE_JS8_10     CreateDependencyResponseRuntime
	NODE_JS10_16    CreateDependencyResponseRuntime
	NODE_JS12_13    CreateDependencyResponseRuntime
	NODE_JS14_18    CreateDependencyResponseRuntime
	PYTHON2_7       CreateDependencyResponseRuntime
	PYTHON3_6       CreateDependencyResponseRuntime
	GO1_8           CreateDependencyResponseRuntime
	GO1_X           CreateDependencyResponseRuntime
	C__NET_CORE_2_0 CreateDependencyResponseRuntime
	C__NET_CORE_2_1 CreateDependencyResponseRuntime
	C__NET_CORE_3_1 CreateDependencyResponseRuntime
	PHP7_3          CreateDependencyResponseRuntime
	PYTHON3_9       CreateDependencyResponseRuntime
}

func GetCreateDependencyResponseRuntimeEnum() CreateDependencyResponseRuntimeEnum {
	return CreateDependencyResponseRuntimeEnum{
		JAVA8: CreateDependencyResponseRuntime{
			value: "Java8",
		},
		JAVA11: CreateDependencyResponseRuntime{
			value: "Java11",
		},
		NODE_JS6_10: CreateDependencyResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: CreateDependencyResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: CreateDependencyResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: CreateDependencyResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: CreateDependencyResponseRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: CreateDependencyResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: CreateDependencyResponseRuntime{
			value: "Python3.6",
		},
		GO1_8: CreateDependencyResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: CreateDependencyResponseRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: CreateDependencyResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: CreateDependencyResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: CreateDependencyResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: CreateDependencyResponseRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: CreateDependencyResponseRuntime{
			value: "Python3.9",
		},
	}
}

func (c CreateDependencyResponseRuntime) Value() string {
	return c.value
}

func (c CreateDependencyResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDependencyResponseRuntime) UnmarshalJSON(b []byte) error {
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
