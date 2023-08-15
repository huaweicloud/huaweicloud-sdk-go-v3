package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFunctionConfigRequestBody struct {

	// 函数名称。
	FuncName string `json:"func_name"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 http: HTTP函数。
	Runtime UpdateFunctionConfigRequestBodyRuntime `json:"runtime"`

	// 函数执行超时时间，超时函数将被强行停止，范围3～900秒，可以通过白名单配置延长到12小时，具体可以咨询华为云函数工作流服务进行配置
	Timeout int32 `json:"timeout"`

	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。
	Handler string `json:"handler"`

	// 函数消耗的内存。 单位M。 取值范围为：128、256、512、768、1024、1280、1536、1792、2048、2560、3072、3584、4096。 最小值为128，最大值为4096。
	MemorySize int32 `json:"memory_size"`

	// 函数消耗的显存，只支持自定义运行时与自定义镜像函数配置GPU。 单位MB。 取值范围为：1024、2048、3072、4096、5120、6144、7168、8192、9216、10240、11264、12288、13312、14336、15360、16384。 最小值为1024，最大值为16384。
	GpuMemory *int32 `json:"gpu_memory,omitempty"`

	// 用户自定义的name/value信息。 在函数中使用的参数。 举例：如函数要访问某个主机，可以设置自定义参数：Host={host_ip}，最多定义20个，总长度不超过4KB。
	UserData *string `json:"user_data,omitempty"`

	// 用户自定义的name/value信息，用于需要加密的配置。
	EncryptedUserData *string `json:"encrypted_user_data,omitempty"`

	// 函数使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	Xrole *string `json:"xrole,omitempty"`

	// 函数app使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	AppXrole *string `json:"app_xrole,omitempty"`

	// 函数描述。
	Description *string `json:"description,omitempty"`

	FuncVpc *FuncVpc `json:"func_vpc,omitempty"`

	MountConfig *MountConfig `json:"mount_config,omitempty"`

	StrategyConfig *StrategyConfig `json:"strategy_config,omitempty"`

	CustomImage *CustomImage `json:"custom_image,omitempty"`

	// 函数扩展配置。
	ExtendConfig *string `json:"extend_config,omitempty"`

	// 函数初始化入口，规则：xx.xx，必须包含“. ”。 举例：对于node.js函数：myfunction.initializer，则表示函数的文件名为myfunction.js，初始化的入口函数名为initializer。
	InitializerHandler *string `json:"initializer_handler,omitempty"`

	// 初始化超时时间，超时函数将被强行停止，范围1～300秒。
	InitializerTimeout *int32 `json:"initializer_timeout,omitempty"`

	// 临时存储大小, 默认512M, 支持配置10G。
	EphemeralStorage *int32 `json:"ephemeral_storage,omitempty"`

	// 企业项目ID，在企业用户创建函数时必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	LogConfig *FuncLogConfig `json:"log_config,omitempty"`

	NetworkController *NetworkControlConfig `json:"network_controller,omitempty"`

	// 是否支持有状态，如果需要支持，需要固定传参为true，v2版本支持
	IsStatefulFunction *bool `json:"is_stateful_function,omitempty"`

	// 是否启动动态内存配置
	EnableDynamicMemory *bool `json:"enable_dynamic_memory,omitempty"`

	// 是否允许在请求头中添加鉴权信息
	EnableAuthInHeader *bool `json:"enable_auth_in_header,omitempty"`

	// 内网域名配置。
	DomainNames *string `json:"domain_names,omitempty"`

	// 函数快照式冷启动Restore Hook入口，仅支持Java，规则：xx.xx，必须包含“. ”。如：com.huawei.demo.Test.restoreHook
	RestoreHookHandler *string `json:"restore_hook_handler,omitempty"`

	// 快照冷启动Restore Hook的超时时间，超时函数将被强行停止，范围1～300秒。
	RestoreHookTimeout *int32 `json:"restore_hook_timeout,omitempty"`
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
	CUSTOM          UpdateFunctionConfigRequestBodyRuntime
	HTTP            UpdateFunctionConfigRequestBodyRuntime
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
		CUSTOM: UpdateFunctionConfigRequestBodyRuntime{
			value: "Custom",
		},
		HTTP: UpdateFunctionConfigRequestBodyRuntime{
			value: "http",
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
