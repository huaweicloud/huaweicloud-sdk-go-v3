package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupabaseApplicationFlavorsRequest Request Object
type ListSupabaseApplicationFlavorsRequest struct {
	XLanguage *ListSupabaseApplicationFlavorsRequestXLanguage `json:"X-Language,omitempty"`

	Version *string `json:"version,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`

	AzCode *string `json:"az_code,omitempty"`

	IsFree *bool `json:"is_free,omitempty"`

	Mode *ListSupabaseApplicationFlavorsRequestMode `json:"mode,omitempty"`
}

func (o ListSupabaseApplicationFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupabaseApplicationFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListSupabaseApplicationFlavorsRequest", string(data)}, " ")
}

type ListSupabaseApplicationFlavorsRequestXLanguage struct {
	value string
}

type ListSupabaseApplicationFlavorsRequestXLanguageEnum struct {
	ZH_CN ListSupabaseApplicationFlavorsRequestXLanguage
	EN_US ListSupabaseApplicationFlavorsRequestXLanguage
}

func GetListSupabaseApplicationFlavorsRequestXLanguageEnum() ListSupabaseApplicationFlavorsRequestXLanguageEnum {
	return ListSupabaseApplicationFlavorsRequestXLanguageEnum{
		ZH_CN: ListSupabaseApplicationFlavorsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSupabaseApplicationFlavorsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSupabaseApplicationFlavorsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSupabaseApplicationFlavorsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupabaseApplicationFlavorsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListSupabaseApplicationFlavorsRequestMode struct {
	value string
}

type ListSupabaseApplicationFlavorsRequestModeEnum struct {
	SINGLE ListSupabaseApplicationFlavorsRequestMode
}

func GetListSupabaseApplicationFlavorsRequestModeEnum() ListSupabaseApplicationFlavorsRequestModeEnum {
	return ListSupabaseApplicationFlavorsRequestModeEnum{
		SINGLE: ListSupabaseApplicationFlavorsRequestMode{
			value: "Single",
		},
	}
}

func (c ListSupabaseApplicationFlavorsRequestMode) Value() string {
	return c.value
}

func (c ListSupabaseApplicationFlavorsRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupabaseApplicationFlavorsRequestMode) UnmarshalJSON(b []byte) error {
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
