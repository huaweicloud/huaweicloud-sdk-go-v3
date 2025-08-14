package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResponseSchemaSubjectDetailsDto 联邦Schema配置Subject属性详细信息
type ResponseSchemaSubjectDetailsDto struct {

	// NameID的格式
	NameIdFormat string `json:"name_id_format"`

	// 是否包含NameID
	Include ResponseSchemaSubjectDetailsDtoInclude `json:"include"`
}

func (o ResponseSchemaSubjectDetailsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseSchemaSubjectDetailsDto struct{}"
	}

	return strings.Join([]string{"ResponseSchemaSubjectDetailsDto", string(data)}, " ")
}

type ResponseSchemaSubjectDetailsDtoInclude struct {
	value string
}

type ResponseSchemaSubjectDetailsDtoIncludeEnum struct {
	REQUIRED ResponseSchemaSubjectDetailsDtoInclude
}

func GetResponseSchemaSubjectDetailsDtoIncludeEnum() ResponseSchemaSubjectDetailsDtoIncludeEnum {
	return ResponseSchemaSubjectDetailsDtoIncludeEnum{
		REQUIRED: ResponseSchemaSubjectDetailsDtoInclude{
			value: "REQUIRED",
		},
	}
}

func (c ResponseSchemaSubjectDetailsDtoInclude) Value() string {
	return c.value
}

func (c ResponseSchemaSubjectDetailsDtoInclude) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseSchemaSubjectDetailsDtoInclude) UnmarshalJSON(b []byte) error {
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
