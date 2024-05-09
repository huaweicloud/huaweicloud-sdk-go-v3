package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFunctionCodeRequestBody struct {

	// 函数代码类型，取值有5种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。 Custom-Image-Swr: 函数代码来源与SWR自定义镜像。
	CodeType UpdateFunctionCodeRequestBodyCodeType `json:"code_type"`

	// 当CodeType为obs时，该值为函数代码包在OBS上的地址，CodeType为其他值时，该字段为空。
	CodeUrl *string `json:"code_url,omitempty"`

	// 函数的文件名，当CodeType为jar/zip时必须提供该字段，inline和obs不需要提供。
	CodeFilename *string `json:"code_filename,omitempty"`

	FuncCode *FuncCode `json:"func_code"`

	// 依赖版本id列表
	DependVersionList *[]string `json:"depend_version_list,omitempty"`
}

func (o UpdateFunctionCodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionCodeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionCodeRequestBody", string(data)}, " ")
}

type UpdateFunctionCodeRequestBodyCodeType struct {
	value string
}

type UpdateFunctionCodeRequestBodyCodeTypeEnum struct {
	INLINE           UpdateFunctionCodeRequestBodyCodeType
	ZIP              UpdateFunctionCodeRequestBodyCodeType
	OBS              UpdateFunctionCodeRequestBodyCodeType
	JAR              UpdateFunctionCodeRequestBodyCodeType
	CUSTOM_IMAGE_SWR UpdateFunctionCodeRequestBodyCodeType
}

func GetUpdateFunctionCodeRequestBodyCodeTypeEnum() UpdateFunctionCodeRequestBodyCodeTypeEnum {
	return UpdateFunctionCodeRequestBodyCodeTypeEnum{
		INLINE: UpdateFunctionCodeRequestBodyCodeType{
			value: "inline",
		},
		ZIP: UpdateFunctionCodeRequestBodyCodeType{
			value: "zip",
		},
		OBS: UpdateFunctionCodeRequestBodyCodeType{
			value: "obs",
		},
		JAR: UpdateFunctionCodeRequestBodyCodeType{
			value: "jar",
		},
		CUSTOM_IMAGE_SWR: UpdateFunctionCodeRequestBodyCodeType{
			value: "Custom-Image-Swr",
		},
	}
}

func (c UpdateFunctionCodeRequestBodyCodeType) Value() string {
	return c.value
}

func (c UpdateFunctionCodeRequestBodyCodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFunctionCodeRequestBodyCodeType) UnmarshalJSON(b []byte) error {
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
