package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type UpdateFunctionCodeResponse struct {

	// 函数的URN（Uniform Resource Name），唯一标识函数。
	FuncUrn *string `json:"func_urn,omitempty"`

	// 函数名称。
	FuncName *string `json:"func_name,omitempty"`

	// 域名id。
	DomainId *string `json:"domain_id,omitempty"`

	// FunctionGraph函数的执行环境 Python2.7: Python语言2.7版本。 Python3.6: Pyton语言3.6版本。 Python3.9: Python语言3.9版本。 Go1.8: Go语言1.8版本。 Go1.x: Go语言1.x版本。 Java8: Java语言8版本。 Java11: Java语言11版本。 Node.js6.10: Nodejs语言6.10版本。 Node.js8.10: Nodejs语言8.10版本。 Node.js10.16: Nodejs语言10.16版本。 Node.js12.13: Nodejs语言12.13版本。 Node.js14.18: Nodejs语言14.18版本。 C#(.NET Core 2.0): C#语言2.0版本。 C#(.NET Core 2.1): C#语言2.1版本。 C#(.NET Core 3.1): C#语言3.1版本。 Custom: 自定义运行时。 PHP7.3: Php语言7.3版本
	Runtime *UpdateFunctionCodeResponseRuntime `json:"runtime,omitempty"`

	// 函数代码类型，取值有4种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。
	CodeType *UpdateFunctionCodeResponseCodeType `json:"code_type,omitempty"`

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

func (o UpdateFunctionCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionCodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateFunctionCodeResponse", string(data)}, " ")
}

type UpdateFunctionCodeResponseRuntime struct {
	value string
}

type UpdateFunctionCodeResponseRuntimeEnum struct {
	JAVA8           UpdateFunctionCodeResponseRuntime
	JAVA11          UpdateFunctionCodeResponseRuntime
	NODE_JS6_10     UpdateFunctionCodeResponseRuntime
	NODE_JS8_10     UpdateFunctionCodeResponseRuntime
	NODE_JS10_16    UpdateFunctionCodeResponseRuntime
	NODE_JS12_13    UpdateFunctionCodeResponseRuntime
	NODE_JS14_18    UpdateFunctionCodeResponseRuntime
	PYTHON2_7       UpdateFunctionCodeResponseRuntime
	PYTHON3_6       UpdateFunctionCodeResponseRuntime
	GO1_8           UpdateFunctionCodeResponseRuntime
	GO1_X           UpdateFunctionCodeResponseRuntime
	C__NET_CORE_2_0 UpdateFunctionCodeResponseRuntime
	C__NET_CORE_2_1 UpdateFunctionCodeResponseRuntime
	C__NET_CORE_3_1 UpdateFunctionCodeResponseRuntime
	PHP7_3          UpdateFunctionCodeResponseRuntime
	PYTHON3_9       UpdateFunctionCodeResponseRuntime
}

func GetUpdateFunctionCodeResponseRuntimeEnum() UpdateFunctionCodeResponseRuntimeEnum {
	return UpdateFunctionCodeResponseRuntimeEnum{
		JAVA8: UpdateFunctionCodeResponseRuntime{
			value: "Java8",
		},
		JAVA11: UpdateFunctionCodeResponseRuntime{
			value: "Java11",
		},
		NODE_JS6_10: UpdateFunctionCodeResponseRuntime{
			value: "Node.js6.10",
		},
		NODE_JS8_10: UpdateFunctionCodeResponseRuntime{
			value: "Node.js8.10",
		},
		NODE_JS10_16: UpdateFunctionCodeResponseRuntime{
			value: "Node.js10.16",
		},
		NODE_JS12_13: UpdateFunctionCodeResponseRuntime{
			value: "Node.js12.13",
		},
		NODE_JS14_18: UpdateFunctionCodeResponseRuntime{
			value: "Node.js14.18",
		},
		PYTHON2_7: UpdateFunctionCodeResponseRuntime{
			value: "Python2.7",
		},
		PYTHON3_6: UpdateFunctionCodeResponseRuntime{
			value: "Python3.6",
		},
		GO1_8: UpdateFunctionCodeResponseRuntime{
			value: "Go1.8",
		},
		GO1_X: UpdateFunctionCodeResponseRuntime{
			value: "Go1.x",
		},
		C__NET_CORE_2_0: UpdateFunctionCodeResponseRuntime{
			value: "C#(.NET Core 2.0)",
		},
		C__NET_CORE_2_1: UpdateFunctionCodeResponseRuntime{
			value: "C#(.NET Core 2.1)",
		},
		C__NET_CORE_3_1: UpdateFunctionCodeResponseRuntime{
			value: "C#(.NET Core 3.1)",
		},
		PHP7_3: UpdateFunctionCodeResponseRuntime{
			value: "PHP7.3",
		},
		PYTHON3_9: UpdateFunctionCodeResponseRuntime{
			value: "Python3.9",
		},
	}
}

func (c UpdateFunctionCodeResponseRuntime) Value() string {
	return c.value
}

func (c UpdateFunctionCodeResponseRuntime) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionCodeResponseRuntime) UnmarshalJSON(b []byte) error {
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

type UpdateFunctionCodeResponseCodeType struct {
	value string
}

type UpdateFunctionCodeResponseCodeTypeEnum struct {
	INLINE UpdateFunctionCodeResponseCodeType
	ZIP    UpdateFunctionCodeResponseCodeType
	OBS    UpdateFunctionCodeResponseCodeType
	JAR    UpdateFunctionCodeResponseCodeType
}

func GetUpdateFunctionCodeResponseCodeTypeEnum() UpdateFunctionCodeResponseCodeTypeEnum {
	return UpdateFunctionCodeResponseCodeTypeEnum{
		INLINE: UpdateFunctionCodeResponseCodeType{
			value: "inline",
		},
		ZIP: UpdateFunctionCodeResponseCodeType{
			value: "zip",
		},
		OBS: UpdateFunctionCodeResponseCodeType{
			value: "obs",
		},
		JAR: UpdateFunctionCodeResponseCodeType{
			value: "jar",
		},
	}
}

func (c UpdateFunctionCodeResponseCodeType) Value() string {
	return c.value
}

func (c UpdateFunctionCodeResponseCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionCodeResponseCodeType) UnmarshalJSON(b []byte) error {
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
