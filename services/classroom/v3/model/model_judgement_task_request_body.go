package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 下发判题任务请求参数
type JudgementTaskRequestBody struct {
	// 第三方指定的判题结果回调url，取值来源于伙伴通道“判题管理配置”-“接口管理”中设置的回调地址相同

	NotifyUrl string `json:"notify_url"`
	// 代码来源：inline（源代码）

	CodeType JudgementTaskRequestBodyCodeType `json:"code_type"`
	// 源代码，需Base64编码

	SourceCode string `json:"source_code"`
	// 任务描述

	Description *string `json:"description,omitempty"`
	// 支持语言类型：java、c、cpp、python

	RuntimeType JudgementTaskRequestBodyRuntimeType `json:"runtime_type"`
	// 代码运行超时时间，单位为秒

	Timeout *int32 `json:"timeout,omitempty"`
	// 结果返回类型：sysout（标准输出）、fileout（以文件形式输出）、imgout（以图片形式输出）、caseout（用例运行返回）、judgeout（输出评判返回）

	OutputType JudgementTaskRequestBodyOutputType `json:"output_type"`
	// 当判题结果类型是caseout和judgeout类型才需要传的字段，表示用例数据

	Testcases *[]JudgementCaseInfo `json:"testcases,omitempty"`
}

func (o JudgementTaskRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JudgementTaskRequestBody struct{}"
	}

	return strings.Join([]string{"JudgementTaskRequestBody", string(data)}, " ")
}

type JudgementTaskRequestBodyCodeType struct {
	value string
}

type JudgementTaskRequestBodyCodeTypeEnum struct {
	INLINE JudgementTaskRequestBodyCodeType
}

func GetJudgementTaskRequestBodyCodeTypeEnum() JudgementTaskRequestBodyCodeTypeEnum {
	return JudgementTaskRequestBodyCodeTypeEnum{
		INLINE: JudgementTaskRequestBodyCodeType{
			value: "inline",
		},
	}
}

func (c JudgementTaskRequestBodyCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *JudgementTaskRequestBodyCodeType) UnmarshalJSON(b []byte) error {
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

type JudgementTaskRequestBodyRuntimeType struct {
	value string
}

type JudgementTaskRequestBodyRuntimeTypeEnum struct {
	JAVA   JudgementTaskRequestBodyRuntimeType
	C      JudgementTaskRequestBodyRuntimeType
	CPP    JudgementTaskRequestBodyRuntimeType
	PYTHON JudgementTaskRequestBodyRuntimeType
}

func GetJudgementTaskRequestBodyRuntimeTypeEnum() JudgementTaskRequestBodyRuntimeTypeEnum {
	return JudgementTaskRequestBodyRuntimeTypeEnum{
		JAVA: JudgementTaskRequestBodyRuntimeType{
			value: "java",
		},
		C: JudgementTaskRequestBodyRuntimeType{
			value: "c",
		},
		CPP: JudgementTaskRequestBodyRuntimeType{
			value: "cpp",
		},
		PYTHON: JudgementTaskRequestBodyRuntimeType{
			value: "python",
		},
	}
}

func (c JudgementTaskRequestBodyRuntimeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *JudgementTaskRequestBodyRuntimeType) UnmarshalJSON(b []byte) error {
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

type JudgementTaskRequestBodyOutputType struct {
	value string
}

type JudgementTaskRequestBodyOutputTypeEnum struct {
	SYSOUT   JudgementTaskRequestBodyOutputType
	FILEOUT  JudgementTaskRequestBodyOutputType
	IMGOUT   JudgementTaskRequestBodyOutputType
	CASEOUT  JudgementTaskRequestBodyOutputType
	JUDGEOUT JudgementTaskRequestBodyOutputType
}

func GetJudgementTaskRequestBodyOutputTypeEnum() JudgementTaskRequestBodyOutputTypeEnum {
	return JudgementTaskRequestBodyOutputTypeEnum{
		SYSOUT: JudgementTaskRequestBodyOutputType{
			value: "sysout",
		},
		FILEOUT: JudgementTaskRequestBodyOutputType{
			value: "fileout",
		},
		IMGOUT: JudgementTaskRequestBodyOutputType{
			value: "imgout",
		},
		CASEOUT: JudgementTaskRequestBodyOutputType{
			value: "caseout",
		},
		JUDGEOUT: JudgementTaskRequestBodyOutputType{
			value: "judgeout",
		},
	}
}

func (c JudgementTaskRequestBodyOutputType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *JudgementTaskRequestBodyOutputType) UnmarshalJSON(b []byte) error {
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
