package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDependencyVersionResponse Response Object
type CreateDependencyVersionResponse struct {

	// 依赖包版本ID。
	Id *string `json:"id,omitempty"`

	// 依赖包拥有者。
	Owner *string `json:"owner,omitempty"`

	// 依赖包在obs的存储地址。
	Link *string `json:"link,omitempty"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 http: HTTP函数。 Custom Image: 自定义镜像函数。
	Runtime *CreateDependencyVersionResponseRuntime `json:"runtime,omitempty"`

	// 依赖包唯一标志。
	Etag *string `json:"etag,omitempty"`

	// 依赖包大小。
	Size *int64 `json:"size,omitempty"`

	// 依赖包名。
	Name *string `json:"name,omitempty"`

	// 依赖包描述。
	Description *string `json:"description,omitempty"`

	// 依赖包文件名。
	FileName *string `json:"file_name,omitempty"`

	// 依赖包版本号
	Version *int64 `json:"version,omitempty"`

	// 依赖包ID
	DepId *string `json:"dep_id,omitempty"`

	// 依赖包更新时间
	LastModified   *int64 `json:"last_modified,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateDependencyVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDependencyVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateDependencyVersionResponse", string(data)}, " ")
}

type CreateDependencyVersionResponseRuntime struct {
	value string
}

type CreateDependencyVersionResponseRuntimeEnum struct {
	JAVA8           CreateDependencyVersionResponseRuntime
	JAVA11          CreateDependencyVersionResponseRuntime
	NODE_JS6_10     CreateDependencyVersionResponseRuntime
	NODE_JS8_10     CreateDependencyVersionResponseRuntime
	NODE_JS10_16    CreateDependencyVersionResponseRuntime
	NODE_JS12_13    CreateDependencyVersionResponseRuntime
	NODE_JS14_18    CreateDependencyVersionResponseRuntime
	PYTHON2_7       CreateDependencyVersionResponseRuntime
	PYTHON3_6       CreateDependencyVersionResponseRuntime
	GO1_8           CreateDependencyVersionResponseRuntime
	GO1_X           CreateDependencyVersionResponseRuntime
	C__NET_CORE_2_0 CreateDependencyVersionResponseRuntime
	C__NET_CORE_2_1 CreateDependencyVersionResponseRuntime
	C__NET_CORE_3_1 CreateDependencyVersionResponseRuntime
	CUSTOM          CreateDependencyVersionResponseRuntime
	PHP7_3          CreateDependencyVersionResponseRuntime
	PYTHON3_9       CreateDependencyVersionResponseRuntime
	HTTP            CreateDependencyVersionResponseRuntime
	CUSTOM_IMAGE    CreateDependencyVersionResponseRuntime
}

func GetCreateDependencyVersionResponseRuntimeEnum() CreateDependencyVersionResponseRuntimeEnum {
	return CreateDependencyVersionResponseRuntimeEnum{
		JAVA8: CreateDependencyVersionResponseRuntime{
			value: "Java8",
		},
		JAVA11: CreateDependencyVersionResponseRuntime{
			value: "Java11",
		},
		NODE_JS6_10: CreateDependencyVersionResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: CreateDependencyVersionResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: CreateDependencyVersionResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: CreateDependencyVersionResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: CreateDependencyVersionResponseRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: CreateDependencyVersionResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: CreateDependencyVersionResponseRuntime{
			value: "Python3.6",
		},
		GO1_8: CreateDependencyVersionResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: CreateDependencyVersionResponseRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: CreateDependencyVersionResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: CreateDependencyVersionResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: CreateDependencyVersionResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		CUSTOM: CreateDependencyVersionResponseRuntime{
			value: "Custom",
		},
		PHP7_3: CreateDependencyVersionResponseRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: CreateDependencyVersionResponseRuntime{
			value: "Python3.9",
		},
		HTTP: CreateDependencyVersionResponseRuntime{
			value: "http",
		},
		CUSTOM_IMAGE: CreateDependencyVersionResponseRuntime{
			value: "Custom Image",
		},
	}
}

func (c CreateDependencyVersionResponseRuntime) Value() string {
	return c.value
}

func (c CreateDependencyVersionResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDependencyVersionResponseRuntime) UnmarshalJSON(b []byte) error {
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
