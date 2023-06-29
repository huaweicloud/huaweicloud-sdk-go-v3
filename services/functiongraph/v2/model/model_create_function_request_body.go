package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateFunctionRequestBody 函数属性结构体。
type CreateFunctionRequestBody struct {

	// 函数名称。
	FuncName string `json:"func_name"`

	// 函数所属的分组Package，用于用户针对函数的自定义分组。
	Package string `json:"package"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 http: HTTP函数。
	Runtime CreateFunctionRequestBodyRuntime `json:"runtime"`

	// 函数执行超时时间，超时函数将被强行停止，范围3～900秒，可以通过白名单配置延长到12小时，具体可以咨询华为云函数工作流服务进行配置
	Timeout int32 `json:"timeout"`

	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。
	Handler string `json:"handler"`

	// 依赖版本id列表
	DependVersionList *[]string `json:"depend_version_list,omitempty"`

	FuncVpc *FuncVpc `json:"func_vpc,omitempty"`

	// 函数消耗的内存。 单位M。 取值范围为：128、256、512、768、1024、1280、1536、1792、2048、2560、3072、3584、4096。 最小值为128，最大值为4096。
	MemorySize int32 `json:"memory_size"`

	// 函数消耗的显存，只支持自定义运行时与自定义镜像函数配置GPU。 单位MB。 取值范围为：1024、2048、3072、4096、5120、6144、7168、8192、9216、10240、11264、12288、13312、14336、15360、16384。 最小值为1024，最大值为16384。
	GpuMemory *int32 `json:"gpu_memory,omitempty"`

	// 函数代码类型，取值有4种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。
	CodeType CreateFunctionRequestBodyCodeType `json:"code_type"`

	// 当CodeType为obs时，该值为函数代码包在OBS上的地址，CodeType为其他值时，该字段为空。
	CodeUrl *string `json:"code_url,omitempty"`

	// 函数的文件名，当CodeType为jar/zip时必须提供该字段，inline和obs不需要提供。
	CodeFilename *string `json:"code_filename,omitempty"`

	// 用户自定义的name/value信息。 在函数中使用的参数。 举例：如函数要访问某个主机，可以设置自定义参数：Host={host_ip}，最多定义20个，总长度不超过4KB。
	UserData *string `json:"user_data,omitempty"`

	// 函数使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	Xrole *string `json:"xrole,omitempty"`

	// 函数app使用的权限委托名称，需要IAM支持，并在IAM界面创建委托，当函数需要访问其他服务时，必须提供该字段。
	AppXrole *string `json:"app_xrole,omitempty"`

	// 函数描述。
	Description *string `json:"description,omitempty"`

	FuncCode *FuncCode `json:"func_code,omitempty"`

	// 函数初始化入口，规则：xx.xx，必须包含“. ”。 举例：对于node.js函数：myfunction.initializer，则表示函数的文件名为myfunction.js，初始化的入口函数名为initializer。
	InitializerHandler *string `json:"initializer_handler,omitempty"`

	// 初始化超时时间，超时函数将被强行停止，范围1～300秒。
	InitializerTimeout *int32 `json:"initializer_timeout,omitempty"`

	// 企业项目ID，在企业用户创建函数时必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 函数版本；部分局点只支持v1函数，缺省值则为v1
	Type *CreateFunctionRequestBodyType `json:"type,omitempty"`

	LogConfig *FuncLogConfig `json:"log_config,omitempty"`

	NetworkController *NetworkControlConfig `json:"network_controller,omitempty"`
}

func (o CreateFunctionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFunctionRequestBody", string(data)}, " ")
}

type CreateFunctionRequestBodyRuntime struct {
	value string
}

type CreateFunctionRequestBodyRuntimeEnum struct {
	JAVA8           CreateFunctionRequestBodyRuntime
	JAVA11          CreateFunctionRequestBodyRuntime
	NODE_JS6_10     CreateFunctionRequestBodyRuntime
	NODE_JS8_10     CreateFunctionRequestBodyRuntime
	NODE_JS10_16    CreateFunctionRequestBodyRuntime
	NODE_JS12_13    CreateFunctionRequestBodyRuntime
	NODE_JS14_18    CreateFunctionRequestBodyRuntime
	PYTHON2_7       CreateFunctionRequestBodyRuntime
	PYTHON3_6       CreateFunctionRequestBodyRuntime
	GO1_8           CreateFunctionRequestBodyRuntime
	GO1_X           CreateFunctionRequestBodyRuntime
	C__NET_CORE_2_0 CreateFunctionRequestBodyRuntime
	C__NET_CORE_2_1 CreateFunctionRequestBodyRuntime
	C__NET_CORE_3_1 CreateFunctionRequestBodyRuntime
	PHP7_3          CreateFunctionRequestBodyRuntime
	PYTHON3_9       CreateFunctionRequestBodyRuntime
	CUSTOM          CreateFunctionRequestBodyRuntime
	HTTP            CreateFunctionRequestBodyRuntime
}

func GetCreateFunctionRequestBodyRuntimeEnum() CreateFunctionRequestBodyRuntimeEnum {
	return CreateFunctionRequestBodyRuntimeEnum{
		JAVA8: CreateFunctionRequestBodyRuntime{
			value: "Java8",
		},
		JAVA11: CreateFunctionRequestBodyRuntime{
			value: "Java11",
		},
		NODE_JS6_10: CreateFunctionRequestBodyRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: CreateFunctionRequestBodyRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: CreateFunctionRequestBodyRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: CreateFunctionRequestBodyRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: CreateFunctionRequestBodyRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: CreateFunctionRequestBodyRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: CreateFunctionRequestBodyRuntime{
			value: "Python3.6",
		},
		GO1_8: CreateFunctionRequestBodyRuntime{
			value: "Go1.8",
		},
		GO1_X: CreateFunctionRequestBodyRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: CreateFunctionRequestBodyRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: CreateFunctionRequestBodyRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: CreateFunctionRequestBodyRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: CreateFunctionRequestBodyRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: CreateFunctionRequestBodyRuntime{
			value: "Python3.9",
		},
		CUSTOM: CreateFunctionRequestBodyRuntime{
			value: "Custom",
		},
		HTTP: CreateFunctionRequestBodyRuntime{
			value: "http",
		},
	}
}

func (c CreateFunctionRequestBodyRuntime) Value() string {
	return c.value
}

func (c CreateFunctionRequestBodyRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionRequestBodyRuntime) UnmarshalJSON(b []byte) error {
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

type CreateFunctionRequestBodyCodeType struct {
	value string
}

type CreateFunctionRequestBodyCodeTypeEnum struct {
	INLINE CreateFunctionRequestBodyCodeType
	ZIP    CreateFunctionRequestBodyCodeType
	OBS    CreateFunctionRequestBodyCodeType
	JAR    CreateFunctionRequestBodyCodeType
}

func GetCreateFunctionRequestBodyCodeTypeEnum() CreateFunctionRequestBodyCodeTypeEnum {
	return CreateFunctionRequestBodyCodeTypeEnum{
		INLINE: CreateFunctionRequestBodyCodeType{
			value: "inline",
		},
		ZIP: CreateFunctionRequestBodyCodeType{
			value: "zip",
		},
		OBS: CreateFunctionRequestBodyCodeType{
			value: "obs",
		},
		JAR: CreateFunctionRequestBodyCodeType{
			value: "jar",
		},
	}
}

func (c CreateFunctionRequestBodyCodeType) Value() string {
	return c.value
}

func (c CreateFunctionRequestBodyCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionRequestBodyCodeType) UnmarshalJSON(b []byte) error {
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

type CreateFunctionRequestBodyType struct {
	value string
}

type CreateFunctionRequestBodyTypeEnum struct {
	V1 CreateFunctionRequestBodyType
	V2 CreateFunctionRequestBodyType
}

func GetCreateFunctionRequestBodyTypeEnum() CreateFunctionRequestBodyTypeEnum {
	return CreateFunctionRequestBodyTypeEnum{
		V1: CreateFunctionRequestBodyType{
			value: "v1",
		},
		V2: CreateFunctionRequestBodyType{
			value: "v2",
		},
	}
}

func (c CreateFunctionRequestBodyType) Value() string {
	return c.value
}

func (c CreateFunctionRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionRequestBodyType) UnmarshalJSON(b []byte) error {
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
