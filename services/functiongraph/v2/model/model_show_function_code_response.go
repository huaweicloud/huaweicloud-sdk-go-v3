package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type ShowFunctionCodeResponse struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FuncUrn *string `json:"func_urn,omitempty"`
	// 函数名称。

	FuncName *string `json:"func_name,omitempty"`
	// 域名id。

	DomainId *string `json:"domain_id,omitempty"`
	// FunctionGraph函数的执行环境 支持Node.js6.10、Python2.7、Python3.6、Java8、Go1.8、Node.js 8.10、C#.NET Core 2.0、C#.NET Core 2.1、PHP7.3。 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本。

	Runtime *ShowFunctionCodeResponseRuntime `json:"runtime,omitempty"`
	// 函数代码类型，取值有4种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。

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
	JAVA_8          ShowFunctionCodeResponseRuntime
	NODE_JS_6_10    ShowFunctionCodeResponseRuntime
	NODE_JS_8_10    ShowFunctionCodeResponseRuntime
	NODE_JS_10_16   ShowFunctionCodeResponseRuntime
	NODE_JS_12_13   ShowFunctionCodeResponseRuntime
	PYTHON_2_7      ShowFunctionCodeResponseRuntime
	PYTHON_3_6      ShowFunctionCodeResponseRuntime
	GO_1_8          ShowFunctionCodeResponseRuntime
	GO_1_X          ShowFunctionCodeResponseRuntime
	C__NET_CORE_2_0 ShowFunctionCodeResponseRuntime
	C__NET_CORE_2_1 ShowFunctionCodeResponseRuntime
	C__NET_CORE_3_1 ShowFunctionCodeResponseRuntime
	PHP_7_3         ShowFunctionCodeResponseRuntime
}

func GetShowFunctionCodeResponseRuntimeEnum() ShowFunctionCodeResponseRuntimeEnum {
	return ShowFunctionCodeResponseRuntimeEnum{
		JAVA_8: ShowFunctionCodeResponseRuntime{
			value: "Java 8",
		},
		NODE_JS_6_10: ShowFunctionCodeResponseRuntime{
			value: "Node.js 6.10",
		},
		NODE_JS_8_10: ShowFunctionCodeResponseRuntime{
			value: "Node.js 8.10",
		},
		NODE_JS_10_16: ShowFunctionCodeResponseRuntime{
			value: "Node.js 10.16",
		},
		NODE_JS_12_13: ShowFunctionCodeResponseRuntime{
			value: "Node.js 12.13",
		},
		PYTHON_2_7: ShowFunctionCodeResponseRuntime{
			value: "Python 2.7",
		},
		PYTHON_3_6: ShowFunctionCodeResponseRuntime{
			value: "Python 3.6",
		},
		GO_1_8: ShowFunctionCodeResponseRuntime{
			value: "Go 1.8",
		},
		GO_1_X: ShowFunctionCodeResponseRuntime{
			value: "Go 1.x",
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
		PHP_7_3: ShowFunctionCodeResponseRuntime{
			value: "PHP 7.3",
		},
	}
}

func (c ShowFunctionCodeResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionCodeResponseRuntime) UnmarshalJSON(b []byte) error {
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

type ShowFunctionCodeResponseCodeType struct {
	value string
}

type ShowFunctionCodeResponseCodeTypeEnum struct {
	INLINE ShowFunctionCodeResponseCodeType
	ZIP    ShowFunctionCodeResponseCodeType
	OBS    ShowFunctionCodeResponseCodeType
	JAR    ShowFunctionCodeResponseCodeType
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
	}
}

func (c ShowFunctionCodeResponseCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionCodeResponseCodeType) UnmarshalJSON(b []byte) error {
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
