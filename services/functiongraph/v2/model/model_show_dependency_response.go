package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowDependencyResponse struct {

	// 依赖包ID。
	Id *string `json:"id,omitempty"`

	// 依赖包拥有者。
	Owner *string `json:"owner,omitempty"`

	// 依赖包在obs的存储地址。
	Link *string `json:"link,omitempty"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本
	Runtime *ShowDependencyResponseRuntime `json:"runtime,omitempty"`

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

func (o ShowDependencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDependencyResponse struct{}"
	}

	return strings.Join([]string{"ShowDependencyResponse", string(data)}, " ")
}

type ShowDependencyResponseRuntime struct {
	value string
}

type ShowDependencyResponseRuntimeEnum struct {
	JAVA8           ShowDependencyResponseRuntime
	JAVA11          ShowDependencyResponseRuntime
	NODE_JS6_10     ShowDependencyResponseRuntime
	NODE_JS8_10     ShowDependencyResponseRuntime
	NODE_JS10_16    ShowDependencyResponseRuntime
	NODE_JS12_13    ShowDependencyResponseRuntime
	NODE_JS14_18    ShowDependencyResponseRuntime
	PYTHON2_7       ShowDependencyResponseRuntime
	PYTHON3_6       ShowDependencyResponseRuntime
	GO1_8           ShowDependencyResponseRuntime
	GO1_X           ShowDependencyResponseRuntime
	C__NET_CORE_2_0 ShowDependencyResponseRuntime
	C__NET_CORE_2_1 ShowDependencyResponseRuntime
	C__NET_CORE_3_1 ShowDependencyResponseRuntime
	PHP7_3          ShowDependencyResponseRuntime
	PYTHON3_9       ShowDependencyResponseRuntime
}

func GetShowDependencyResponseRuntimeEnum() ShowDependencyResponseRuntimeEnum {
	return ShowDependencyResponseRuntimeEnum{
		JAVA8: ShowDependencyResponseRuntime{
			value: "Java8",
		},
		JAVA11: ShowDependencyResponseRuntime{
			value: "Java11",
		},
		NODE_JS6_10: ShowDependencyResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: ShowDependencyResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: ShowDependencyResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: ShowDependencyResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: ShowDependencyResponseRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: ShowDependencyResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: ShowDependencyResponseRuntime{
			value: "Python3.6",
		},
		GO1_8: ShowDependencyResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: ShowDependencyResponseRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: ShowDependencyResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: ShowDependencyResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: ShowDependencyResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: ShowDependencyResponseRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: ShowDependencyResponseRuntime{
			value: "Python3.9",
		},
	}
}

func (c ShowDependencyResponseRuntime) Value() string {
	return c.value
}

func (c ShowDependencyResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDependencyResponseRuntime) UnmarshalJSON(b []byte) error {
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
