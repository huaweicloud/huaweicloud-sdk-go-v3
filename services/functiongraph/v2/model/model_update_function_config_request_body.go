package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFunctionConfigRequestBody struct {

	// 函数名称。
	FuncName string `json:"func_name"`

	// FunctionGraph函数的执行环境 Java8: Java语言8版本。 Java11: Java语言11版本。 Java17: Java语言17版本（当前仅支持华北-乌兰察布二零二） Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Python3.10: Python语言3.10版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 Node.js16.17: Nodejs语言16.17版本。 Node.js18.15: Nodejs语言18.15版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 C#(.NET Core 6.0): C#语言6.0版本（当前仅支持华北-乌兰察布二零二）。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 Cangjie1.0：仓颉语言1.0版本。 http: HTTP函数。 Custom Image: 自定义镜像函数。
	Runtime UpdateFunctionConfigRequestBodyRuntime `json:"runtime"`

	// 函数执行超时时间，超时函数将被强行停止，范围3～259200秒。
	Timeout int32 `json:"timeout"`

	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。
	Handler string `json:"handler"`

	// 函数消耗的内存。 单位M。 取值范围为：128、256、512、768、1024、1280、1536、1792、2048、2560、3072、3584、4096。 最小值为128，最大值为4096。
	MemorySize int32 `json:"memory_size"`

	// 函数消耗的显存，只支持自定义运行时与自定义镜像函数配置GPU。 单位MB。 取值范围为：1024、2048、3072、4096、5120、6144、7168、8192、9216、10240、11264、12288、13312、14336、15360、16384。 最小值为1024，最大值为16384。
	GpuMemory *int32 `json:"gpu_memory,omitempty"`

	// 显卡类型。
	GpuType *string `json:"gpu_type,omitempty"`

	// 用户自定义的name/value信息。 在函数中使用的参数。 举例：如函数要访问某个主机，可以设置自定义参数：Host={host_ip}，最多定义20个，总长度不超过4KB。
	UserData *string `json:"user_data,omitempty"`

	// 用户自定义的name/value信息，用于需要加密的配置。
	EncryptedUserData *string `json:"encrypted_user_data,omitempty"`

	// 函数配置委托。需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。配置后用户可以通过函数执行入口方法中的context参数获取具有委托中权限的token、ak、sk，用于访问其他云服务。如果用户函数不访问任何云服务，则不用提供委托名称。
	Xrole *string `json:"xrole,omitempty"`

	// 函数执行委托。可为函数执行单独配置执行委托，这将减小不必要的性能损耗；不单独配置执行委托时，函数执行和函数配置将使用同一委托。
	AppXrole *string `json:"app_xrole,omitempty"`

	// 函数描述。
	Description *string `json:"description,omitempty"`

	FuncVpc *FuncVpc `json:"func_vpc,omitempty"`

	MountConfig *MountConfig `json:"mount_config,omitempty"`

	StrategyConfig *StrategyConfig `json:"strategy_config,omitempty"`

	CustomImage *CustomImage `json:"custom_image,omitempty"`

	// 函数扩展配置。
	ExtendConfig *string `json:"extend_config,omitempty"`

	// 函数初始化入口，规则：xx.xx，必须包含“. ”。当配置初始化函数时，此参数必填。 举例：对于node.js函数：myfunction.initializer，则表示函数的文件名为myfunction.js，初始化的入口函数名为initializer。
	InitializerHandler *string `json:"initializer_handler,omitempty"`

	// 初始化超时时间，超时函数将被强行停止，范围1～300秒。当配置初始化函数时，此参数必填。
	InitializerTimeout *int32 `json:"initializer_timeout,omitempty"`

	// 函数预停止函数的入口，规则：xx.xx，必须包含“. ”。 举例：对于node.js函数：myfunction.pre_stop_handler，则表示函数的文件名为myfunction.js，初始化的入口函数名为pre_stop_handler。
	PreStopHandler *string `json:"pre_stop_handler,omitempty"`

	// 初始化超时时间，超时函数将被强行停止，范围1～90秒。
	PreStopTimeout *int32 `json:"pre_stop_timeout,omitempty"`

	// 临时存储大小。默认情况下会为函数的/tmp目录分配512MB的空间。您可以通过临时存储设置将函数的/tmp目录大小调整为10G。
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

	// 函数快照式冷启动Restore Hook入口，仅支持Java，规则：xx.xx，必须包含“. ”。如：com.xxx.demo.Test.restoreHook
	RestoreHookHandler *string `json:"restore_hook_handler,omitempty"`

	// 快照冷启动Restore Hook的超时时间，超时函数将被强行停止，范围1～300秒。
	RestoreHookTimeout *int32 `json:"restore_hook_timeout,omitempty"`

	// 心跳函数函数的入口，规则：xx.xx，必须包含“. ”，只支持JAVA运行时配置。 心跳函数入口需要与函数执行入口在同一文件下。在开启心跳函数配置时，此参数必填。
	HeartbeatHandler *string `json:"heartbeat_handler,omitempty"`

	// 类隔离开关，只支持JAVA运行时配置。开启类隔离后可以支持Kafka转储并提升类加载效率，但也可能会导致某些兼容性问题，请谨慎开启。
	EnableClassIsolation *bool `json:"enable_class_isolation,omitempty"`
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
	JAVA17          UpdateFunctionConfigRequestBodyRuntime
	PYTHON2_7       UpdateFunctionConfigRequestBodyRuntime
	PYTHON3_6       UpdateFunctionConfigRequestBodyRuntime
	PYTHON3_9       UpdateFunctionConfigRequestBodyRuntime
	PYTHON3_10      UpdateFunctionConfigRequestBodyRuntime
	GO1_8           UpdateFunctionConfigRequestBodyRuntime
	GO1_X           UpdateFunctionConfigRequestBodyRuntime
	NODE_JS6_10     UpdateFunctionConfigRequestBodyRuntime
	NODE_JS8_10     UpdateFunctionConfigRequestBodyRuntime
	NODE_JS10_16    UpdateFunctionConfigRequestBodyRuntime
	NODE_JS12_13    UpdateFunctionConfigRequestBodyRuntime
	NODE_JS14_18    UpdateFunctionConfigRequestBodyRuntime
	NODE_JS16_17    UpdateFunctionConfigRequestBodyRuntime
	NODE_JS18_15    UpdateFunctionConfigRequestBodyRuntime
	C__NET_CORE_2_0 UpdateFunctionConfigRequestBodyRuntime
	C__NET_CORE_2_1 UpdateFunctionConfigRequestBodyRuntime
	C__NET_CORE_3_1 UpdateFunctionConfigRequestBodyRuntime
	C__NET_CORE_6_0 UpdateFunctionConfigRequestBodyRuntime
	CUSTOM          UpdateFunctionConfigRequestBodyRuntime
	PHP7_3          UpdateFunctionConfigRequestBodyRuntime
	CANGJIE1_0      UpdateFunctionConfigRequestBodyRuntime
	HTTP            UpdateFunctionConfigRequestBodyRuntime
	CUSTOM_IMAGE    UpdateFunctionConfigRequestBodyRuntime
}

func GetUpdateFunctionConfigRequestBodyRuntimeEnum() UpdateFunctionConfigRequestBodyRuntimeEnum {
	return UpdateFunctionConfigRequestBodyRuntimeEnum{
		JAVA8: UpdateFunctionConfigRequestBodyRuntime{
			value: "Java8",
		},
		JAVA11: UpdateFunctionConfigRequestBodyRuntime{
			value: "Java11",
		},
		JAVA17: UpdateFunctionConfigRequestBodyRuntime{
			value: "Java17",
		},
		PYTHON2_7: UpdateFunctionConfigRequestBodyRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: UpdateFunctionConfigRequestBodyRuntime{
			value: "Python3.6",
		},
		PYTHON3_9: UpdateFunctionConfigRequestBodyRuntime{
			value: "Python3.9",
		},
		PYTHON3_10: UpdateFunctionConfigRequestBodyRuntime{
			value: "Python3.10",
		},
		GO1_8: UpdateFunctionConfigRequestBodyRuntime{
			value: "Go1.8",
		},
		GO1_X: UpdateFunctionConfigRequestBodyRuntime{
			value: "Go1.x",
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
		NODE_JS16_17: UpdateFunctionConfigRequestBodyRuntime{
			value: "Node.js16.17",
		},
		NODE_JS18_15: UpdateFunctionConfigRequestBodyRuntime{
			value: "Node.js18.15",
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
		C__NET_CORE_6_0: UpdateFunctionConfigRequestBodyRuntime{
			value: "C#(.NET Core 6.0)",
		},
		CUSTOM: UpdateFunctionConfigRequestBodyRuntime{
			value: "Custom",
		},
		PHP7_3: UpdateFunctionConfigRequestBodyRuntime{
			value: "PHP7.3",
		},
		CANGJIE1_0: UpdateFunctionConfigRequestBodyRuntime{
			value: "Cangjie1.0",
		},
		HTTP: UpdateFunctionConfigRequestBodyRuntime{
			value: "http",
		},
		CUSTOM_IMAGE: UpdateFunctionConfigRequestBodyRuntime{
			value: "Custom Image",
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
