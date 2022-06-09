package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type CreateFunctionVersionResponse struct {

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

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本
	Runtime *CreateFunctionVersionResponseRuntime `json:"runtime,omitempty"`

	// 函数执行超时时间，超时函数将被强行停止，范围3～900秒
	Timeout *int32 `json:"timeout,omitempty"`

	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。
	Handler *string `json:"handler,omitempty"`

	// 函数消耗的内存。 单位M。 取值范围为：128、256、512、768、1024、1280、1536、1792、2048、2560、3072、3584、4096。 最小值为128，最大值为4096。
	MemorySize *int32 `json:"memory_size,omitempty"`

	// 函数占用的cpu资源。 单位为millicore（1 core=1000 millicores）。 取值与MemorySize成比例，默认是128M内存占0.1个核（100 millicores）。 函数占用的CPU为基础CPU：200 millicores，再加上内存按比例占用的CPU，计算方法：内存/128 *100 + 200。
	Cpu *int32 `json:"cpu,omitempty"`

	// 函数代码类型，取值有4种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。
	CodeType *CreateFunctionVersionResponseCodeType `json:"code_type,omitempty"`

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

	// 函数使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	Xrole *string `json:"xrole,omitempty"`

	// 函数app使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	AppXrole *string `json:"app_xrole,omitempty"`

	// 函数描述。
	Description *string `json:"description,omitempty"`

	// 函数版本描述。
	VersionDescription *string `json:"version_description,omitempty"`

	// 函数最后一次更新时间。
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty"`

	FuncVpc *FuncVpc `json:"func_vpc,omitempty"`

	MountConfig *MountConfig `json:"mount_config,omitempty"`

	StrategyConfig *StrategyConfig `json:"strategy_config,omitempty"`

	// 函数依赖代码包列表。
	Dependencies *[]Dependency `json:"dependencies,omitempty"`

	// 函数初始化入口，规则：xx.xx，必须包含“. ”。 举例：对于node.js函数：myfunction.initializer，则表示函数的文件名为myfunction.js，初始化的入口函数名为initializer。
	InitializerHandler *string `json:"initializer_handler,omitempty"`

	// 初始化超时时间，超时函数将被强行停止，范围1～300秒。
	InitializerTimeout *int32 `json:"initializer_timeout,omitempty"`

	// 企业项目ID，在企业用户创建函数时必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否允许进行长时间超时设置。
	LongTime *bool `json:"long_time,omitempty"`

	// 自定义日志查询组id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 自定义日志查询流id
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// v2表示为公测版本,v1为原来版本。
	Type *CreateFunctionVersionResponseType `json:"type,omitempty"`

	// 是否允许动态内存配置
	EnableDynamicMemory *bool `json:"enable_dynamic_memory,omitempty"`

	FunctionAsyncConfig *FunctionAsyncConfig `json:"function_async_config,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o CreateFunctionVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateFunctionVersionResponse", string(data)}, " ")
}

type CreateFunctionVersionResponseRuntime struct {
	value string
}

type CreateFunctionVersionResponseRuntimeEnum struct {
	JAVA8           CreateFunctionVersionResponseRuntime
	JAVA11          CreateFunctionVersionResponseRuntime
	NODE_JS6_10     CreateFunctionVersionResponseRuntime
	NODE_JS8_10     CreateFunctionVersionResponseRuntime
	NODE_JS10_16    CreateFunctionVersionResponseRuntime
	NODE_JS12_13    CreateFunctionVersionResponseRuntime
	NODE_JS14_18    CreateFunctionVersionResponseRuntime
	PYTHON2_7       CreateFunctionVersionResponseRuntime
	PYTHON3_6       CreateFunctionVersionResponseRuntime
	GO1_8           CreateFunctionVersionResponseRuntime
	GO1_X           CreateFunctionVersionResponseRuntime
	C__NET_CORE_2_0 CreateFunctionVersionResponseRuntime
	C__NET_CORE_2_1 CreateFunctionVersionResponseRuntime
	C__NET_CORE_3_1 CreateFunctionVersionResponseRuntime
	PHP7_3          CreateFunctionVersionResponseRuntime
	PYTHON3_9       CreateFunctionVersionResponseRuntime
}

func GetCreateFunctionVersionResponseRuntimeEnum() CreateFunctionVersionResponseRuntimeEnum {
	return CreateFunctionVersionResponseRuntimeEnum{
		JAVA8: CreateFunctionVersionResponseRuntime{
			value: "Java8",
		},
		JAVA11: CreateFunctionVersionResponseRuntime{
			value: "Java11",
		},
		NODE_JS6_10: CreateFunctionVersionResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: CreateFunctionVersionResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: CreateFunctionVersionResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: CreateFunctionVersionResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: CreateFunctionVersionResponseRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: CreateFunctionVersionResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: CreateFunctionVersionResponseRuntime{
			value: "Python3.6",
		},
		GO1_8: CreateFunctionVersionResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: CreateFunctionVersionResponseRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: CreateFunctionVersionResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: CreateFunctionVersionResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: CreateFunctionVersionResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: CreateFunctionVersionResponseRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: CreateFunctionVersionResponseRuntime{
			value: "Python3.9",
		},
	}
}

func (c CreateFunctionVersionResponseRuntime) Value() string {
	return c.value
}

func (c CreateFunctionVersionResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionVersionResponseRuntime) UnmarshalJSON(b []byte) error {
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

type CreateFunctionVersionResponseCodeType struct {
	value string
}

type CreateFunctionVersionResponseCodeTypeEnum struct {
	INLINE CreateFunctionVersionResponseCodeType
	ZIP    CreateFunctionVersionResponseCodeType
	OBS    CreateFunctionVersionResponseCodeType
	JAR    CreateFunctionVersionResponseCodeType
}

func GetCreateFunctionVersionResponseCodeTypeEnum() CreateFunctionVersionResponseCodeTypeEnum {
	return CreateFunctionVersionResponseCodeTypeEnum{
		INLINE: CreateFunctionVersionResponseCodeType{
			value: "inline",
		},
		ZIP: CreateFunctionVersionResponseCodeType{
			value: "zip",
		},
		OBS: CreateFunctionVersionResponseCodeType{
			value: "obs",
		},
		JAR: CreateFunctionVersionResponseCodeType{
			value: "jar",
		},
	}
}

func (c CreateFunctionVersionResponseCodeType) Value() string {
	return c.value
}

func (c CreateFunctionVersionResponseCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionVersionResponseCodeType) UnmarshalJSON(b []byte) error {
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

type CreateFunctionVersionResponseType struct {
	value string
}

type CreateFunctionVersionResponseTypeEnum struct {
	V1 CreateFunctionVersionResponseType
	V2 CreateFunctionVersionResponseType
}

func GetCreateFunctionVersionResponseTypeEnum() CreateFunctionVersionResponseTypeEnum {
	return CreateFunctionVersionResponseTypeEnum{
		V1: CreateFunctionVersionResponseType{
			value: "v1",
		},
		V2: CreateFunctionVersionResponseType{
			value: "v2",
		},
	}
}

func (c CreateFunctionVersionResponseType) Value() string {
	return c.value
}

func (c CreateFunctionVersionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionVersionResponseType) UnmarshalJSON(b []byte) error {
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
