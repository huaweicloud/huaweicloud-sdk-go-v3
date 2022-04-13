package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// Response Object
type ImportFunctionResponse struct {
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
	// FunctionGraph函数的执行环境 支持Node.js6.10、Python2.7、Python3.6、Java8、Go1.8、Node.js 8.10、C#.NET Core 2.0、C#.NET Core 2.1、PHP7.3。 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 Java11、Nodejs14:、Python3:在type为v2时支持

	Runtime *ImportFunctionResponseRuntime `json:"runtime,omitempty"`
	// 函数执行超时时间，超时函数将被强行停止，范围3～900秒

	Timeout *int32 `json:"timeout,omitempty"`
	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。

	Handler *string `json:"handler,omitempty"`
	// 函数消耗的内存。 单位M。 取值范围为：128、256、512、768、1024、1280、1536、1792、2048、2560、3072、3584、4096。 最小值为128，最大值为4096。

	MemorySize *int32 `json:"memory_size,omitempty"`
	// 函数占用的cpu资源。 单位为millicore（1 core=1000 millicores）。 取值与MemorySize成比例，默认是128M内存占0.1个核（100 millicores）。 函数占用的CPU为基础CPU：200 millicores，再加上内存按比例占用的CPU，计算方法：内存/128 *100 + 200。

	Cpu *int32 `json:"cpu,omitempty"`
	// 函数代码类型，取值有4种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。

	CodeType *ImportFunctionResponseCodeType `json:"code_type,omitempty"`
	// 当CodeType为obs时，该值为函数代码包在OBS上的地址，CodeType为其他值时，该字段为空。

	CodeUrl *string `json:"code_url,omitempty"`
	// 函数的文件名，当CodeType为jar/zip时必须提供该字段，inline和obs不需要提供。

	CodeFilename *string `json:"code_filename,omitempty"`
	// 函数大小，单位：字节。

	CodeSize *int64 `json:"code_size,omitempty"`
	// 用户自定义的name/value信息。 在函数中使用的参数。 举例：如函数要访问某个主机，可以设置自定义参数：Host={host_ip}，最多定义20个，总长度不超过4KB。

	UserData *string `json:"user_data,omitempty"`
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
	// 依赖id列表

	DependList *[]string `json:"depend_list,omitempty"`

	StrategyConfig *StrategyConfig `json:"strategy_config,omitempty"`
	// 函数扩展配置。

	ExtendConfig *string `json:"extend_config,omitempty"`
	// 函数初始化入口，规则：xx.xx，必须包含“. ”。 举例：对于node.js函数：myfunction.initializer，则表示函数的文件名为myfunction.js，初始化的入口函数名为initializer。

	InitializerHandler *string `json:"initializer_handler,omitempty"`
	// 初始化超时时间，超时函数将被强行停止，范围1～300秒。

	InitializerTimeout *int32 `json:"initializer_timeout,omitempty"`
	// 企业项目ID，在企业用户创建函数时必填。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ImportFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFunctionResponse struct{}"
	}

	return strings.Join([]string{"ImportFunctionResponse", string(data)}, " ")
}

type ImportFunctionResponseRuntime struct {
	value string
}

type ImportFunctionResponseRuntimeEnum struct {
	JAVA8           ImportFunctionResponseRuntime
	JAVA11          ImportFunctionResponseRuntime
	NODE_JS6_10     ImportFunctionResponseRuntime
	NODE_JS8_10     ImportFunctionResponseRuntime
	NODE_JS10_16    ImportFunctionResponseRuntime
	NODE_JS12_13    ImportFunctionResponseRuntime
	NODE_JS14_18    ImportFunctionResponseRuntime
	PYTHON2_7       ImportFunctionResponseRuntime
	PYTHON3_6       ImportFunctionResponseRuntime
	PYTHON3_9       ImportFunctionResponseRuntime
	GO1_8           ImportFunctionResponseRuntime
	GO1_X           ImportFunctionResponseRuntime
	C__NET_CORE_2_0 ImportFunctionResponseRuntime
	C__NET_CORE_2_1 ImportFunctionResponseRuntime
	C__NET_CORE_3_1 ImportFunctionResponseRuntime
	PHP7_3          ImportFunctionResponseRuntime
}

func GetImportFunctionResponseRuntimeEnum() ImportFunctionResponseRuntimeEnum {
	return ImportFunctionResponseRuntimeEnum{
		JAVA8: ImportFunctionResponseRuntime{
			value: "Java8",
		},
		JAVA11: ImportFunctionResponseRuntime{
			value: "Java11",
		},
		NODE_JS6_10: ImportFunctionResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: ImportFunctionResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: ImportFunctionResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: ImportFunctionResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: ImportFunctionResponseRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: ImportFunctionResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: ImportFunctionResponseRuntime{
			value: "Python3.6",
		},
		PYTHON3_9: ImportFunctionResponseRuntime{
			value: "Python3.9",
		},
		GO1_8: ImportFunctionResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: ImportFunctionResponseRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: ImportFunctionResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: ImportFunctionResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: ImportFunctionResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: ImportFunctionResponseRuntime{
			value: "PHP7.3",
		},
	}
}

func (c ImportFunctionResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportFunctionResponseRuntime) UnmarshalJSON(b []byte) error {
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

type ImportFunctionResponseCodeType struct {
	value string
}

type ImportFunctionResponseCodeTypeEnum struct {
	INLINE ImportFunctionResponseCodeType
	ZIP    ImportFunctionResponseCodeType
	OBS    ImportFunctionResponseCodeType
	JAR    ImportFunctionResponseCodeType
}

func GetImportFunctionResponseCodeTypeEnum() ImportFunctionResponseCodeTypeEnum {
	return ImportFunctionResponseCodeTypeEnum{
		INLINE: ImportFunctionResponseCodeType{
			value: "inline",
		},
		ZIP: ImportFunctionResponseCodeType{
			value: "zip",
		},
		OBS: ImportFunctionResponseCodeType{
			value: "obs",
		},
		JAR: ImportFunctionResponseCodeType{
			value: "jar",
		},
	}
}

func (c ImportFunctionResponseCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportFunctionResponseCodeType) UnmarshalJSON(b []byte) error {
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
