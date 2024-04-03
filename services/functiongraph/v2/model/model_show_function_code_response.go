package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowFunctionCodeResponse Response Object
type ShowFunctionCodeResponse struct {

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FuncUrn *string `json:"func_urn,omitempty"`

	// 函数名称。
	FuncName *string `json:"func_name,omitempty"`

	// 域名id。
	DomainId *string `json:"domain_id,omitempty"`

	// FunctionGraph函数的执行环境 Java8: Java语言8版本。 Java11: Java语言11版本。 Java17: Java语言17版本（当前仅支持华北-乌兰察布二零二） Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Python3.10: Python语言3.10版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 Node.js16.17: Nodejs语言16.17版本。 Node.js18.15: Nodejs语言18.15版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 C#(.NET Core 6.0): C#语言6.0版本（当前仅支持华北-乌兰察布二零二）。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。 Cangjie1.0：仓颉语言1.0版本。 http: HTTP函数。 Custom Image: 自定义镜像函数。
	Runtime *ShowFunctionCodeResponseRuntime `json:"runtime,omitempty"`

	// 函数代码类型，取值有5种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。 Custom-Image-Swr: 函数代码来源与SWR自定义镜像。
	CodeType *ShowFunctionCodeResponseCodeType `json:"code_type,omitempty"`

	// 当CodeType为obs时，该值为函数代码包在OBS上的地址，CodeType为其他值时，该字段为空。
	CodeUrl *string `json:"code_url,omitempty"`

	// 函数的文件名，当CodeType为jar/zip时必须提供该字段，inline和obs不需要提供。
	CodeFilename *string `json:"code_filename,omitempty"`

	// 函数大小，单位：字节。
	CodeSize *int64 `json:"code_size,omitempty"`

	// 函数代码SHA512 hash值，用于判断函数是否变化。
	Digest *string `json:"digest,omitempty"`

	// 函数最后一次更新时间。
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty"`

	FuncCode *FuncCode `json:"func_code,omitempty"`

	// 依赖id列表
	DependList *[]string `json:"depend_list,omitempty"`

	// 依赖版本id列表
	DependVersionList *[]string `json:"depend_version_list,omitempty"`

	StrategyConfig *StrategyConfig `json:"strategy_config,omitempty"`

	// 函数依赖代码包列表。
	Dependencies   *[]Dependency `json:"dependencies,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowFunctionCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionCodeResponse struct{}"
	}

	return strings.Join([]string{"ShowFunctionCodeResponse", string(data)}, " ")
}

type ShowFunctionCodeResponseRuntime struct {
	value string
}

type ShowFunctionCodeResponseRuntimeEnum struct {
	JAVA8           ShowFunctionCodeResponseRuntime
	JAVA11          ShowFunctionCodeResponseRuntime
	JAVA17          ShowFunctionCodeResponseRuntime
	PYTHON2_7       ShowFunctionCodeResponseRuntime
	PYTHON3_6       ShowFunctionCodeResponseRuntime
	PYTHON3_9       ShowFunctionCodeResponseRuntime
	PYTHON3_10      ShowFunctionCodeResponseRuntime
	GO1_8           ShowFunctionCodeResponseRuntime
	GO1_X           ShowFunctionCodeResponseRuntime
	NODE_JS6_10     ShowFunctionCodeResponseRuntime
	NODE_JS8_10     ShowFunctionCodeResponseRuntime
	NODE_JS10_16    ShowFunctionCodeResponseRuntime
	NODE_JS12_13    ShowFunctionCodeResponseRuntime
	NODE_JS14_18    ShowFunctionCodeResponseRuntime
	NODE_JS16_17    ShowFunctionCodeResponseRuntime
	NODE_JS18_15    ShowFunctionCodeResponseRuntime
	C__NET_CORE_2_0 ShowFunctionCodeResponseRuntime
	C__NET_CORE_2_1 ShowFunctionCodeResponseRuntime
	C__NET_CORE_3_1 ShowFunctionCodeResponseRuntime
	C__NET_CORE_6_0 ShowFunctionCodeResponseRuntime
	CUSTOM          ShowFunctionCodeResponseRuntime
	PHP7_3          ShowFunctionCodeResponseRuntime
	CANGJIE1_0      ShowFunctionCodeResponseRuntime
	HTTP            ShowFunctionCodeResponseRuntime
	CUSTOM_IMAGE    ShowFunctionCodeResponseRuntime
}

func GetShowFunctionCodeResponseRuntimeEnum() ShowFunctionCodeResponseRuntimeEnum {
	return ShowFunctionCodeResponseRuntimeEnum{
		JAVA8: ShowFunctionCodeResponseRuntime{
			value: "Java8",
		},
		JAVA11: ShowFunctionCodeResponseRuntime{
			value: "Java11",
		},
		JAVA17: ShowFunctionCodeResponseRuntime{
			value: "Java17",
		},
		PYTHON2_7: ShowFunctionCodeResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: ShowFunctionCodeResponseRuntime{
			value: "Python3.6",
		},
		PYTHON3_9: ShowFunctionCodeResponseRuntime{
			value: "Python3.9",
		},
		PYTHON3_10: ShowFunctionCodeResponseRuntime{
			value: "Python3.10",
		},
		GO1_8: ShowFunctionCodeResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: ShowFunctionCodeResponseRuntime{
			value: "Go1.x",
		},
		NODE_JS6_10: ShowFunctionCodeResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: ShowFunctionCodeResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: ShowFunctionCodeResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: ShowFunctionCodeResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: ShowFunctionCodeResponseRuntime{
			value: "Node.js14.18",
		},
		NODE_JS16_17: ShowFunctionCodeResponseRuntime{
			value: "Node.js16.17",
		},
		NODE_JS18_15: ShowFunctionCodeResponseRuntime{
			value: "Node.js18.15",
		},
		C__NET_CORE_2_0: ShowFunctionCodeResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: ShowFunctionCodeResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: ShowFunctionCodeResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		C__NET_CORE_6_0: ShowFunctionCodeResponseRuntime{
			value: "C#(.NET Core 6.0)",
		},
		CUSTOM: ShowFunctionCodeResponseRuntime{
			value: "Custom",
		},
		PHP7_3: ShowFunctionCodeResponseRuntime{
			value: "PHP7.3",
		},
		CANGJIE1_0: ShowFunctionCodeResponseRuntime{
			value: "Cangjie1.0",
		},
		HTTP: ShowFunctionCodeResponseRuntime{
			value: "http",
		},
		CUSTOM_IMAGE: ShowFunctionCodeResponseRuntime{
			value: "Custom Image",
		},
	}
}

func (c ShowFunctionCodeResponseRuntime) Value() string {
	return c.value
}

func (c ShowFunctionCodeResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionCodeResponseRuntime) UnmarshalJSON(b []byte) error {
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

type ShowFunctionCodeResponseCodeType struct {
	value string
}

type ShowFunctionCodeResponseCodeTypeEnum struct {
	INLINE           ShowFunctionCodeResponseCodeType
	ZIP              ShowFunctionCodeResponseCodeType
	OBS              ShowFunctionCodeResponseCodeType
	JAR              ShowFunctionCodeResponseCodeType
	CUSTOM_IMAGE_SWR ShowFunctionCodeResponseCodeType
}

func GetShowFunctionCodeResponseCodeTypeEnum() ShowFunctionCodeResponseCodeTypeEnum {
	return ShowFunctionCodeResponseCodeTypeEnum{
		INLINE: ShowFunctionCodeResponseCodeType{
			value: "inline",
		},
		ZIP: ShowFunctionCodeResponseCodeType{
			value: "zip",
		},
		OBS: ShowFunctionCodeResponseCodeType{
			value: "obs",
		},
		JAR: ShowFunctionCodeResponseCodeType{
			value: "jar",
		},
		CUSTOM_IMAGE_SWR: ShowFunctionCodeResponseCodeType{
			value: "Custom-Image-Swr",
		},
	}
}

func (c ShowFunctionCodeResponseCodeType) Value() string {
	return c.value
}

func (c ShowFunctionCodeResponseCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionCodeResponseCodeType) UnmarshalJSON(b []byte) error {
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
