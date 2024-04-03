package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateFunctionConfigResponse Response Object
type UpdateFunctionConfigResponse struct {

	// 函数id，唯一标识函数。
	FuncId *string `json:"func_id,omitempty"`

	// 函数资源id。
	ResourceId *string `json:"resource_id,omitempty"`

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FuncUrn *string `json:"func_urn,omitempty"`

	// 函数名称。
	FuncName *string `json:"func_name,omitempty"`

	// 域名id。
	DomainId *string `json:"domain_id,omitempty"`

	// 租户的project id。
	Namespace *string `json:"namespace,omitempty"`

	// 租户的project name。
	ProjectName *string `json:"project_name,omitempty"`

	// 函数所属的分组Package，用于用户针对函数的自定义分组。
	Package *string `json:"package,omitempty"`

	// FunctionGraph函数的执行环境 Java8: Java语言8版本。 Java11: Java语言11版本。 Java17: Java语言17版本（当前仅支持华北-乌兰察布二零二） Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Python3.10: Python语言3.10版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 Node.js16.17: Nodejs语言16.17版本。 Node.js18.15: Nodejs语言18.15版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 C#(.NET Core 6.0): C#语言6.0版本（当前仅支持华北-乌兰察布二零二）。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 Cangjie1.0：仓颉语言1.0版本。 http: HTTP函数。 Custom Image: 自定义镜像函数。
	Runtime *UpdateFunctionConfigResponseRuntime `json:"runtime,omitempty"`

	// 函数执行超时时间，超时函数将被强行停止，范围3～259200秒。
	Timeout *int32 `json:"timeout,omitempty"`

	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。
	Handler *string `json:"handler,omitempty"`

	// 函数消耗的内存。 单位M。 取值范围为：128、256、512、768、1024、1280、1536、1792、2048、2560、3072、3584、4096。 最小值为128，最大值为4096。
	MemorySize *int32 `json:"memory_size,omitempty"`

	// 函数消耗的显存，只支持自定义运行时与自定义镜像函数配置GPU。 单位MB。 取值范围为：1024、2048、3072、4096、5120、6144、7168、8192、9216、10240、11264、12288、13312、14336、15360、16384。 最小值为1024，最大值为16384。
	GpuMemory *int32 `json:"gpu_memory,omitempty"`

	// 函数占用的cpu资源。 单位为millicore（1 core=1000 millicores）。 取值与MemorySize成比例，默认是128M内存占0.1个核（100 millicores）。
	Cpu *int32 `json:"cpu,omitempty"`

	// 函数代码类型，取值有5种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。 Custom-Image-Swr: 函数代码来源与SWR自定义镜像。
	CodeType *UpdateFunctionConfigResponseCodeType `json:"code_type,omitempty"`

	// 当CodeType为obs时，该值为函数代码包在OBS上的地址，CodeType为其他值时，该字段为空。
	CodeUrl *string `json:"code_url,omitempty"`

	// 函数的文件名，当CodeType为jar/zip时必须提供该字段，inline和obs不需要提供。
	CodeFilename *string `json:"code_filename,omitempty"`

	// 函数大小，单位：字节。
	CodeSize *int64 `json:"code_size,omitempty"`

	// 用户自定义的name/value信息。 在函数中使用的参数。 举例：如函数要访问某个主机，可以设置自定义参数：Host={host_ip}，最多定义20个，总长度不超过4KB。
	UserData *string `json:"user_data,omitempty"`

	// 用户自定义的name/value信息，用于需要加密的配置。
	EncryptedUserData *string `json:"encrypted_user_data,omitempty"`

	// 函数代码SHA512 hash值，用于判断函数是否变化。
	Digest *string `json:"digest,omitempty"`

	// 函数版本号，由系统自动生成，规则：vYYYYMMDD-HHMMSS（v+年月日-时分秒）。
	Version *string `json:"version,omitempty"`

	// 函数版本的内部标识。
	ImageName *string `json:"image_name,omitempty"`

	// 函数配置委托。需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。配置后用户可以通过函数执行入口方法中的context参数获取具有委托中权限的token、ak、sk，用于访问其他云服务。如果用户函数不访问任何云服务，则不用提供委托名称。
	Xrole *string `json:"xrole,omitempty"`

	// 函数执行委托。可为函数执行单独配置执行委托，这将减小不必要的性能损耗；不单独配置执行委托时，函数执行和函数配置将使用同一委托。
	AppXrole *string `json:"app_xrole,omitempty"`

	// 函数描述。
	Description *string `json:"description,omitempty"`

	// 函数最后一次更新时间。
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty"`

	// 临时存储大小。默认情况下会为函数的/tmp目录分配512MB的空间。您可以通过临时存储设置将函数的/tmp目录大小调整为10G。
	EphemeralStorage *int32 `json:"ephemeral_storage,omitempty"`

	FuncVpc *FuncVpc `json:"func_vpc,omitempty"`

	MountConfig *MountConfig `json:"mount_config,omitempty"`

	StrategyConfig *StrategyConfig `json:"strategy_config,omitempty"`

	// 函数依赖代码包列表。
	Dependencies *[]Dependency `json:"dependencies,omitempty"`

	// 函数初始化入口，规则：xx.xx，必须包含“. ”。当配置初始化函数时，此参数必填。 举例：对于node.js函数：myfunction.initializer，则表示函数的文件名为myfunction.js，初始化的入口函数名为initializer。
	InitializerHandler *string `json:"initializer_handler,omitempty"`

	// 初始化超时时间，超时函数将被强行停止，范围1～300秒。当配置初始化函数时，此参数必填。
	InitializerTimeout *int32 `json:"initializer_timeout,omitempty"`

	// 函数预停止函数的入口，规则：xx.xx，必须包含“. ”。 举例：对于node.js函数：myfunction.pre_stop_handler，则表示函数的文件名为myfunction.js，初始化的入口函数名为pre_stop_handler。
	PreStopHandler *string `json:"pre_stop_handler,omitempty"`

	// 初始化超时时间，超时函数将被强行停止，范围1～90秒。
	PreStopTimeout *int32 `json:"pre_stop_timeout,omitempty"`

	// 企业项目ID，在企业用户创建函数时必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否允许进行长时间超时设置。
	LongTime *bool `json:"long_time,omitempty"`

	// 自定义日志查询组id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 自定义日志查询流id
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// v2表示为正式版本,v1为废弃版本。
	Type *UpdateFunctionConfigResponseType `json:"type,omitempty"`

	// 适配CloudDebug场景，是否开启云调试（已废弃）
	EnableCloudDebug *string `json:"enable_cloud_debug,omitempty"`

	// 是否启动动态内存配置
	EnableDynamicMemory *bool `json:"enable_dynamic_memory,omitempty"`

	// 是否支持有状态，v2版本支持
	IsStatefulFunction *bool `json:"is_stateful_function,omitempty"`

	// 函数配置的需要支持域名解析的内网域名。
	DomainNames *string `json:"domain_names,omitempty"`

	// 是否允许在请求头中添加鉴权信息，只支持自定义镜像函数
	EnableAuthInHeader *bool `json:"enable_auth_in_header,omitempty"`

	CustomImage *CustomImage `json:"custom_image,omitempty"`

	// 是否返回流式数据（已废弃）
	IsReturnStream *bool `json:"is_return_stream,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateFunctionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateFunctionConfigResponse", string(data)}, " ")
}

type UpdateFunctionConfigResponseRuntime struct {
	value string
}

type UpdateFunctionConfigResponseRuntimeEnum struct {
	JAVA8           UpdateFunctionConfigResponseRuntime
	JAVA11          UpdateFunctionConfigResponseRuntime
	JAVA17          UpdateFunctionConfigResponseRuntime
	PYTHON2_7       UpdateFunctionConfigResponseRuntime
	PYTHON3_6       UpdateFunctionConfigResponseRuntime
	PYTHON3_9       UpdateFunctionConfigResponseRuntime
	PYTHON3_10      UpdateFunctionConfigResponseRuntime
	GO1_8           UpdateFunctionConfigResponseRuntime
	GO1_X           UpdateFunctionConfigResponseRuntime
	NODE_JS6_10     UpdateFunctionConfigResponseRuntime
	NODE_JS8_10     UpdateFunctionConfigResponseRuntime
	NODE_JS10_16    UpdateFunctionConfigResponseRuntime
	NODE_JS12_13    UpdateFunctionConfigResponseRuntime
	NODE_JS14_18    UpdateFunctionConfigResponseRuntime
	NODE_JS16_17    UpdateFunctionConfigResponseRuntime
	NODE_JS18_15    UpdateFunctionConfigResponseRuntime
	C__NET_CORE_2_0 UpdateFunctionConfigResponseRuntime
	C__NET_CORE_2_1 UpdateFunctionConfigResponseRuntime
	C__NET_CORE_3_1 UpdateFunctionConfigResponseRuntime
	C__NET_CORE_6_0 UpdateFunctionConfigResponseRuntime
	CUSTOM          UpdateFunctionConfigResponseRuntime
	PHP7_3          UpdateFunctionConfigResponseRuntime
	CANGJIE1_0      UpdateFunctionConfigResponseRuntime
	HTTP            UpdateFunctionConfigResponseRuntime
	CUSTOM_IMAGE    UpdateFunctionConfigResponseRuntime
}

func GetUpdateFunctionConfigResponseRuntimeEnum() UpdateFunctionConfigResponseRuntimeEnum {
	return UpdateFunctionConfigResponseRuntimeEnum{
		JAVA8: UpdateFunctionConfigResponseRuntime{
			value: "Java8",
		},
		JAVA11: UpdateFunctionConfigResponseRuntime{
			value: "Java11",
		},
		JAVA17: UpdateFunctionConfigResponseRuntime{
			value: "Java17",
		},
		PYTHON2_7: UpdateFunctionConfigResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: UpdateFunctionConfigResponseRuntime{
			value: "Python3.6",
		},
		PYTHON3_9: UpdateFunctionConfigResponseRuntime{
			value: "Python3.9",
		},
		PYTHON3_10: UpdateFunctionConfigResponseRuntime{
			value: "Python3.10",
		},
		GO1_8: UpdateFunctionConfigResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: UpdateFunctionConfigResponseRuntime{
			value: "Go1.x",
		},
		NODE_JS6_10: UpdateFunctionConfigResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: UpdateFunctionConfigResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: UpdateFunctionConfigResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: UpdateFunctionConfigResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: UpdateFunctionConfigResponseRuntime{
			value: "Node.js14.18",
		},
		NODE_JS16_17: UpdateFunctionConfigResponseRuntime{
			value: "Node.js16.17",
		},
		NODE_JS18_15: UpdateFunctionConfigResponseRuntime{
			value: "Node.js18.15",
		},
		C__NET_CORE_2_0: UpdateFunctionConfigResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: UpdateFunctionConfigResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: UpdateFunctionConfigResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		C__NET_CORE_6_0: UpdateFunctionConfigResponseRuntime{
			value: "C#(.NET Core 6.0)",
		},
		CUSTOM: UpdateFunctionConfigResponseRuntime{
			value: "Custom",
		},
		PHP7_3: UpdateFunctionConfigResponseRuntime{
			value: "PHP7.3",
		},
		CANGJIE1_0: UpdateFunctionConfigResponseRuntime{
			value: "Cangjie1.0",
		},
		HTTP: UpdateFunctionConfigResponseRuntime{
			value: "http",
		},
		CUSTOM_IMAGE: UpdateFunctionConfigResponseRuntime{
			value: "Custom Image",
		},
	}
}

func (c UpdateFunctionConfigResponseRuntime) Value() string {
	return c.value
}

func (c UpdateFunctionConfigResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionConfigResponseRuntime) UnmarshalJSON(b []byte) error {
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

type UpdateFunctionConfigResponseCodeType struct {
	value string
}

type UpdateFunctionConfigResponseCodeTypeEnum struct {
	INLINE           UpdateFunctionConfigResponseCodeType
	ZIP              UpdateFunctionConfigResponseCodeType
	OBS              UpdateFunctionConfigResponseCodeType
	JAR              UpdateFunctionConfigResponseCodeType
	CUSTOM_IMAGE_SWR UpdateFunctionConfigResponseCodeType
}

func GetUpdateFunctionConfigResponseCodeTypeEnum() UpdateFunctionConfigResponseCodeTypeEnum {
	return UpdateFunctionConfigResponseCodeTypeEnum{
		INLINE: UpdateFunctionConfigResponseCodeType{
			value: "inline",
		},
		ZIP: UpdateFunctionConfigResponseCodeType{
			value: "zip",
		},
		OBS: UpdateFunctionConfigResponseCodeType{
			value: "obs",
		},
		JAR: UpdateFunctionConfigResponseCodeType{
			value: "jar",
		},
		CUSTOM_IMAGE_SWR: UpdateFunctionConfigResponseCodeType{
			value: "Custom-Image-Swr",
		},
	}
}

func (c UpdateFunctionConfigResponseCodeType) Value() string {
	return c.value
}

func (c UpdateFunctionConfigResponseCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionConfigResponseCodeType) UnmarshalJSON(b []byte) error {
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

type UpdateFunctionConfigResponseType struct {
	value string
}

type UpdateFunctionConfigResponseTypeEnum struct {
	V1 UpdateFunctionConfigResponseType
	V2 UpdateFunctionConfigResponseType
}

func GetUpdateFunctionConfigResponseTypeEnum() UpdateFunctionConfigResponseTypeEnum {
	return UpdateFunctionConfigResponseTypeEnum{
		V1: UpdateFunctionConfigResponseType{
			value: "v1",
		},
		V2: UpdateFunctionConfigResponseType{
			value: "v2",
		},
	}
}

func (c UpdateFunctionConfigResponseType) Value() string {
	return c.value
}

func (c UpdateFunctionConfigResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionConfigResponseType) UnmarshalJSON(b []byte) error {
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
