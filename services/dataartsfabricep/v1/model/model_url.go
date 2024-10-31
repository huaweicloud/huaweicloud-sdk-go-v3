package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Url Url信息
type Url struct {

	// Url名称
	Name *string `json:"name,omitempty"`

	// Url地址
	Address *string `json:"address,omitempty"`

	// 类型，PUBLIC为公网地址，PRIVATE为内网地址
	Type *UrlType `json:"type,omitempty"`
}

func (o Url) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Url struct{}"
	}

	return strings.Join([]string{"Url", string(data)}, " ")
}

type UrlType struct {
	value string
}

type UrlTypeEnum struct {
	PRIVATE UrlType
	PUBLIC  UrlType
}

func GetUrlTypeEnum() UrlTypeEnum {
	return UrlTypeEnum{
		PRIVATE: UrlType{
			value: "PRIVATE",
		},
		PUBLIC: UrlType{
			value: "PUBLIC",
		},
	}
}

func (c UrlType) Value() string {
	return c.value
}

func (c UrlType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlType) UnmarshalJSON(b []byte) error {
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
