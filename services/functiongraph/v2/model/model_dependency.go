package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Dependency 函数依赖包结构。
type Dependency struct {

	// 依赖包属主的domainId。
	Owner string `json:"owner"`

	// 依赖包在OBS上的链接。
	Link string `json:"link"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 http: HTTP函数。
	Runtime DependencyRuntime `json:"runtime"`

	// 依赖包的md5值
	Etag string `json:"etag"`

	// 依赖包大小。
	Size int64 `json:"size"`

	// 依赖包名称。
	Name string `json:"name"`

	// 依赖包描述。
	Description string `json:"description"`

	// 依赖包文件名，如果创建方式为zip时。
	FileName *string `json:"file_name,omitempty"`
}

func (o Dependency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dependency struct{}"
	}

	return strings.Join([]string{"Dependency", string(data)}, " ")
}

type DependencyRuntime struct {
	value string
}

type DependencyRuntimeEnum struct {
	JAVA8           DependencyRuntime
	JAVA11          DependencyRuntime
	NODE_JS6_10     DependencyRuntime
	NODE_JS8_10     DependencyRuntime
	NODE_JS10_16    DependencyRuntime
	NODE_JS12_13    DependencyRuntime
	NODE_JS14_18    DependencyRuntime
	PYTHON2_7       DependencyRuntime
	PYTHON3_6       DependencyRuntime
	GO1_8           DependencyRuntime
	GO1_X           DependencyRuntime
	C__NET_CORE_2_0 DependencyRuntime
	C__NET_CORE_2_1 DependencyRuntime
	C__NET_CORE_3_1 DependencyRuntime
	PHP7_3          DependencyRuntime
	PYTHON3_9       DependencyRuntime
	CUSTOM          DependencyRuntime
	HTTP            DependencyRuntime
}

func GetDependencyRuntimeEnum() DependencyRuntimeEnum {
	return DependencyRuntimeEnum{
		JAVA8: DependencyRuntime{
			value: "Java8",
		},
		JAVA11: DependencyRuntime{
			value: "Java11",
		},
		NODE_JS6_10: DependencyRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: DependencyRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: DependencyRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: DependencyRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: DependencyRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: DependencyRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: DependencyRuntime{
			value: "Python3.6",
		},
		GO1_8: DependencyRuntime{
			value: "Go1.8",
		},
		GO1_X: DependencyRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: DependencyRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: DependencyRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: DependencyRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: DependencyRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: DependencyRuntime{
			value: "Python3.9",
		},
		CUSTOM: DependencyRuntime{
			value: "Custom",
		},
		HTTP: DependencyRuntime{
			value: "http",
		},
	}
}

func (c DependencyRuntime) Value() string {
	return c.value
}

func (c DependencyRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DependencyRuntime) UnmarshalJSON(b []byte) error {
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
