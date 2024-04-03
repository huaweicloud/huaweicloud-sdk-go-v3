package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListDependenciesResult struct {

	// 依赖包ID
	Id string `json:"id"`

	// 依赖包拥有者，public标识为公共依赖包
	Owner string `json:"owner"`

	// 依赖包在obs的存储地址
	Link string `json:"link"`

	// FunctionGraph函数的执行环境 Java8: Java语言8版本。 Java11: Java语言11版本。 Java17: Java语言17版本（当前仅支持华北-乌兰察布二零二） Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Python3.10: Python语言3.10版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 Node.js16.17: Nodejs语言16.17版本。 Node.js18.15: Nodejs语言18.15版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 C#(.NET Core 6.0): C#语言6.0版本（当前仅支持华北-乌兰察布二零二）。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 Cangjie1.0：仓颉语言1.0版本。 http: HTTP函数。 Custom Image: 自定义镜像函数。
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
	JAVA17          ListDependenciesResultRuntime
	PYTHON2_7       ListDependenciesResultRuntime
	PYTHON3_6       ListDependenciesResultRuntime
	PYTHON3_9       ListDependenciesResultRuntime
	PYTHON3_10      ListDependenciesResultRuntime
	GO1_8           ListDependenciesResultRuntime
	GO1_X           ListDependenciesResultRuntime
	NODE_JS6_10     ListDependenciesResultRuntime
	NODE_JS8_10     ListDependenciesResultRuntime
	NODE_JS10_16    ListDependenciesResultRuntime
	NODE_JS12_13    ListDependenciesResultRuntime
	NODE_JS14_18    ListDependenciesResultRuntime
	NODE_JS16_17    ListDependenciesResultRuntime
	NODE_JS18_15    ListDependenciesResultRuntime
	C__NET_CORE_2_0 ListDependenciesResultRuntime
	C__NET_CORE_2_1 ListDependenciesResultRuntime
	C__NET_CORE_3_1 ListDependenciesResultRuntime
	C__NET_CORE_6_0 ListDependenciesResultRuntime
	CUSTOM          ListDependenciesResultRuntime
	PHP7_3          ListDependenciesResultRuntime
	CANGJIE1_0      ListDependenciesResultRuntime
	HTTP            ListDependenciesResultRuntime
	CUSTOM_IMAGE    ListDependenciesResultRuntime
}

func GetListDependenciesResultRuntimeEnum() ListDependenciesResultRuntimeEnum {
	return ListDependenciesResultRuntimeEnum{
		JAVA8: ListDependenciesResultRuntime{
			value: "Java8",
		},
		JAVA11: ListDependenciesResultRuntime{
			value: "Java11",
		},
		JAVA17: ListDependenciesResultRuntime{
			value: "Java17",
		},
		PYTHON2_7: ListDependenciesResultRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: ListDependenciesResultRuntime{
			value: "Python3.6",
		},
		PYTHON3_9: ListDependenciesResultRuntime{
			value: "Python3.9",
		},
		PYTHON3_10: ListDependenciesResultRuntime{
			value: "Python3.10",
		},
		GO1_8: ListDependenciesResultRuntime{
			value: "Go1.8",
		},
		GO1_X: ListDependenciesResultRuntime{
			value: "Go1.x",
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
		NODE_JS16_17: ListDependenciesResultRuntime{
			value: "Node.js16.17",
		},
		NODE_JS18_15: ListDependenciesResultRuntime{
			value: "Node.js18.15",
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
		C__NET_CORE_6_0: ListDependenciesResultRuntime{
			value: "C#(.NET Core 6.0)",
		},
		CUSTOM: ListDependenciesResultRuntime{
			value: "Custom",
		},
		PHP7_3: ListDependenciesResultRuntime{
			value: "PHP7.3",
		},
		CANGJIE1_0: ListDependenciesResultRuntime{
			value: "Cangjie1.0",
		},
		HTTP: ListDependenciesResultRuntime{
			value: "http",
		},
		CUSTOM_IMAGE: ListDependenciesResultRuntime{
			value: "Custom Image",
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
