package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MaterialComponentInfo 可替换的素材组件信息
type MaterialComponentInfo struct {

	// 素材组件名称。
	ComponentName string `json:"component_name"`

	// 素材组件类型。 * CLOTHES：衣服 * PANTS：裤子 * SHOES：鞋子 * HAIR：头发 * EYELASH：睫毛 * EYEBROW：眉毛
	ComponentType MaterialComponentInfoComponentType `json:"component_type"`

	// 素材组件描述。
	ComponentDesc *string `json:"component_desc,omitempty"`
}

func (o MaterialComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaterialComponentInfo struct{}"
	}

	return strings.Join([]string{"MaterialComponentInfo", string(data)}, " ")
}

type MaterialComponentInfoComponentType struct {
	value string
}

type MaterialComponentInfoComponentTypeEnum struct {
	CLOTHES MaterialComponentInfoComponentType
	PANTS   MaterialComponentInfoComponentType
	SHOES   MaterialComponentInfoComponentType
	HAIR    MaterialComponentInfoComponentType
	EYELASH MaterialComponentInfoComponentType
	EYEBROW MaterialComponentInfoComponentType
}

func GetMaterialComponentInfoComponentTypeEnum() MaterialComponentInfoComponentTypeEnum {
	return MaterialComponentInfoComponentTypeEnum{
		CLOTHES: MaterialComponentInfoComponentType{
			value: "CLOTHES",
		},
		PANTS: MaterialComponentInfoComponentType{
			value: "PANTS",
		},
		SHOES: MaterialComponentInfoComponentType{
			value: "SHOES",
		},
		HAIR: MaterialComponentInfoComponentType{
			value: "HAIR",
		},
		EYELASH: MaterialComponentInfoComponentType{
			value: "EYELASH",
		},
		EYEBROW: MaterialComponentInfoComponentType{
			value: "EYEBROW",
		},
	}
}

func (c MaterialComponentInfoComponentType) Value() string {
	return c.value
}

func (c MaterialComponentInfoComponentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MaterialComponentInfoComponentType) UnmarshalJSON(b []byte) error {
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
