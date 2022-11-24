package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomizeSchemaUpdateReq struct {

	// 事件模型描述
	Description *string `json:"description,omitempty"`

	// 事件模型兼容性
	Compatibility *CustomizeSchemaUpdateReqCompatibility `json:"compatibility,omitempty"`
}

func (o CustomizeSchemaUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSchemaUpdateReq struct{}"
	}

	return strings.Join([]string{"CustomizeSchemaUpdateReq", string(data)}, " ")
}

type CustomizeSchemaUpdateReqCompatibility struct {
	value string
}

type CustomizeSchemaUpdateReqCompatibilityEnum struct {
	NONE     CustomizeSchemaUpdateReqCompatibility
	FORWARD  CustomizeSchemaUpdateReqCompatibility
	BACKWARD CustomizeSchemaUpdateReqCompatibility
	FULL     CustomizeSchemaUpdateReqCompatibility
}

func GetCustomizeSchemaUpdateReqCompatibilityEnum() CustomizeSchemaUpdateReqCompatibilityEnum {
	return CustomizeSchemaUpdateReqCompatibilityEnum{
		NONE: CustomizeSchemaUpdateReqCompatibility{
			value: "NONE",
		},
		FORWARD: CustomizeSchemaUpdateReqCompatibility{
			value: "FORWARD",
		},
		BACKWARD: CustomizeSchemaUpdateReqCompatibility{
			value: "BACKWARD",
		},
		FULL: CustomizeSchemaUpdateReqCompatibility{
			value: "FULL",
		},
	}
}

func (c CustomizeSchemaUpdateReqCompatibility) Value() string {
	return c.value
}

func (c CustomizeSchemaUpdateReqCompatibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSchemaUpdateReqCompatibility) UnmarshalJSON(b []byte) error {
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
