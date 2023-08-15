package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListDependenciesResult struct {

	// 依赖包ID
	Id string `json:"id"`

	// 依赖包拥有者，public标识为公共依赖包
	Owner string `json:"owner"`

	// 依赖包在obs的存储地址
	Link string `json:"link"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 http: HTTP函数。
	Runtime ListDependenciesResultRuntime `json:"runtime"`

	// 依赖包唯一标志
	Etag string `json:"etag"`

	// 依赖包大小
	Size int64 `json:"size"`

	// 依赖包名
	Name string `json:"name"`

	// 依赖包文件名
	FileName *string `json:"file_name,omitempty"`

	// 依赖包描述。
	Description *string `json:"description,omitempty"`

	// 依赖包版本号
	Version *int64 `json:"version,omitempty"`

	// 依赖包更新时间
	LastModified *int64 `json:"last_modified,omitempty"`
}

func (o ListDependenciesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDependenciesResult struct{}"
	}

	return strings.Join([]string{"ListDependenciesResult", string(data)}, " ")
}

type ListDependenciesResultRuntime struct {
	value string
}

type ListDependenciesResultRuntimeEnum struct {
	JAVA8           ListDependenciesResultRuntime
	JAVA11          ListDependenciesResultRuntime
	NODE_JS6_10     ListDependenciesResultRuntime
	NODE_JS8_10     ListDependenciesResultRuntime
	NODE_JS10_16    ListDependenciesResultRuntime
	NODE_JS12_13    ListDependenciesResultRuntime
	NODE_JS14_18    ListDependenciesResultRuntime
	PYTHON2_7       ListDependenciesResultRuntime
	PYTHON3_6       ListDependenciesResultRuntime
	GO1_8           ListDependenciesResultRuntime
	GO1_X           ListDependenciesResultRuntime
	C__NET_CORE_2_0 ListDependenciesResultRuntime
	C__NET_CORE_2_1 ListDependenciesResultRuntime
	C__NET_CORE_3_1 ListDependenciesResultRuntime
	PHP7_3          ListDependenciesResultRuntime
	PYTHON3_9       ListDependenciesResultRuntime
	CUSTOM          ListDependenciesResultRuntime
	HTTP            ListDependenciesResultRuntime
}

func GetListDependenciesResultRuntimeEnum() ListDependenciesResultRuntimeEnum {
	return ListDependenciesResultRuntimeEnum{
		JAVA8: ListDependenciesResultRuntime{
			value: "Java8",
		},
		JAVA11: ListDependenciesResultRuntime{
			value: "Java11",
		},
		NODE_JS6_10: ListDependenciesResultRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: ListDependenciesResultRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: ListDependenciesResultRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: ListDependenciesResultRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: ListDependenciesResultRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: ListDependenciesResultRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: ListDependenciesResultRuntime{
			value: "Python3.6",
		},
		GO1_8: ListDependenciesResultRuntime{
			value: "Go1.8",
		},
		GO1_X: ListDependenciesResultRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: ListDependenciesResultRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: ListDependenciesResultRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: ListDependenciesResultRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: ListDependenciesResultRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: ListDependenciesResultRuntime{
			value: "Python3.9",
		},
		CUSTOM: ListDependenciesResultRuntime{
			value: "Custom",
		},
		HTTP: ListDependenciesResultRuntime{
			value: "http",
		},
	}
}

func (c ListDependenciesResultRuntime) Value() string {
	return c.value
}

func (c ListDependenciesResultRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDependenciesResultRuntime) UnmarshalJSON(b []byte) error {
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
