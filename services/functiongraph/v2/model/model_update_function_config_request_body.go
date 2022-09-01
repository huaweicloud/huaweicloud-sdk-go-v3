package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFunctionConfigRequestBody struct {

	// 函数名称。
	FuncName string `json:"func_name" xml:"func_name"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本
	Runtime UpdateFunctionConfigRequestBodyRuntime `json:"runtime" xml:"runtime"`

	// 函数执行超时时间，超时函数将被强行停止，范围3～900秒，可以通过白名单配置延长到12小时，具体可以咨询华为云函数工作流服务进行配置
	Timeout int32 `json:"timeout" xml:"timeout"`

	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。
	Handler string `json:"handler" xml:"handler"`

	// 函数消耗的内存。 单位M。 取值范围为：128、256、512、768、1024、1280、1536、1792、2048、2560、3072、3584、4096。 最小值为128，最大值为4096。
	MemorySize int32 `json:"memory_size" xml:"memory_size"`

	// 用户自定义的name/value信息。 在函数中使用的参数。 举例：如函数要访问某个主机，可以设置自定义参数：Host={host_ip}，最多定义20个，总长度不超过4KB。
	UserData *string `json:"user_data,omitempty" xml:"user_data"`

	// 用户自定义的name/value信息，用于需要加密的配置。
	EncryptedUserData *string `json:"encrypted_user_data,omitempty" xml:"encrypted_user_data"`

	// 函数使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	Xrole *string `json:"xrole,omitempty" xml:"xrole"`

	// 函数app使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	AppXrole *string `json:"app_xrole,omitempty" xml:"app_xrole"`

	// 函数描述。
	Description *string `json:"description,omitempty" xml:"description"`

	FuncVpc *FuncVpc `json:"func_vpc,omitempty" xml:"func_vpc"`

	MountConfig *MountConfig `json:"mount_config,omitempty" xml:"mount_config"`

	StrategyConfig *StrategyConfig `json:"strategy_config,omitempty" xml:"strategy_config"`

	// 函数扩展配置。
	ExtendConfig *string `json:"extend_config,omitempty" xml:"extend_config"`

	// 函数初始化入口，规则：xx.xx，必须包含“. ”。 举例：对于node.js函数：myfunction.initializer，则表示函数的文件名为myfunction.js，初始化的入口函数名为initializer。
	InitializerHandler *string `json:"initializer_handler,omitempty" xml:"initializer_handler"`

	// 初始化超时时间，超时函数将被强行停止，范围1～300秒。
	InitializerTimeout *int32 `json:"initializer_timeout,omitempty" xml:"initializer_timeout"`

	// 企业项目ID，在企业用户创建函数时必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 是否支持有状态，如果需要支持，需要固定传参为true，v2版本支持
	IsStatefulFunction *bool `json:"is_stateful_function,omitempty" xml:"is_stateful_function"`

	// 内网域名配置。
	DomainNames *string `json:"domain_names,omitempty" xml:"domain_names"`
}

func (o UpdateFunctionConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionConfigRequestBody", string(data)}, " ")
}

type UpdateFunctionConfigRequestBodyRuntime struct {
	value string
}

type UpdateFunctionConfigRequestBodyRuntimeEnum struct {
	JAVA8           UpdateFunctionConfigRequestBodyRuntime
	JAVA11          UpdateFunctionConfigRequestBodyRuntime
	NODE_JS6_10     UpdateFunctionConfigRequestBodyRuntime
	NODE_JS8_10     UpdateFunctionConfigRequestBodyRuntime
	NODE_JS10_16    UpdateFunctionConfigRequestBodyRuntime
	NODE_JS12_13    UpdateFunctionConfigRequestBodyRuntime
	NODE_JS14_18    UpdateFunctionConfigRequestBodyRuntime
	PYTHON2_7       UpdateFunctionConfigRequestBodyRuntime
	PYTHON3_6       UpdateFunctionConfigRequestBodyRuntime
	GO1_8           UpdateFunctionConfigRequestBodyRuntime
	GO1_X           UpdateFunctionConfigRequestBodyRuntime
	C__NET_CORE_2_0 UpdateFunctionConfigRequestBodyRuntime
	C__NET_CORE_2_1 UpdateFunctionConfigRequestBodyRuntime
	C__NET_CORE_3_1 UpdateFunctionConfigRequestBodyRuntime
	PHP7_3          UpdateFunctionConfigRequestBodyRuntime
	PYTHON3_9       UpdateFunctionConfigRequestBodyRuntime
}

func GetUpdateFunctionConfigRequestBodyRuntimeEnum() UpdateFunctionConfigRequestBodyRuntimeEnum {
	return UpdateFunctionConfigRequestBodyRuntimeEnum{
		JAVA8: UpdateFunctionConfigRequestBodyRuntime{
			value: "Java8",
		},
		JAVA11: UpdateFunctionConfigRequestBodyRuntime{
			value: "Java11",
		},
		NODE_JS6_10: UpdateFunctionConfigRequestBodyRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: UpdateFunctionConfigRequestBodyRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: UpdateFunctionConfigRequestBodyRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: UpdateFunctionConfigRequestBodyRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: UpdateFunctionConfigRequestBodyRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: UpdateFunctionConfigRequestBodyRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: UpdateFunctionConfigRequestBodyRuntime{
			value: "Python3.6",
		},
		GO1_8: UpdateFunctionConfigRequestBodyRuntime{
			value: "Go1.8",
		},
		GO1_X: UpdateFunctionConfigRequestBodyRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: UpdateFunctionConfigRequestBodyRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: UpdateFunctionConfigRequestBodyRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: UpdateFunctionConfigRequestBodyRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: UpdateFunctionConfigRequestBodyRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: UpdateFunctionConfigRequestBodyRuntime{
			value: "Python3.9",
		},
	}
}

func (c UpdateFunctionConfigRequestBodyRuntime) Value() string {
	return c.value
}

func (c UpdateFunctionConfigRequestBodyRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionConfigRequestBodyRuntime) UnmarshalJSON(b []byte) error {
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
