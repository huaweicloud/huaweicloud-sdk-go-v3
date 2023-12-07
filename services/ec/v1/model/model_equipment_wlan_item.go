package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EquipmentWlanItem struct {

	// 是否支持wlan，提供给UI识别是否显示Wi-Fi配置页面
	SupportWlan bool `json:"support_wlan"`

	// 是否使能wlan，取值为true时，必须填写name、security_enabled、name_hided
	WlanEnabled bool `json:"wlan_enabled"`

	// Wi-Fi名称，长度1-32个字符，不支持中文字符，特殊字符只支持!~@_.?
	Name *string `json:"name,omitempty"`

	// 是否开启无线安全，取值为true时，必须填写authentication_method、encrption_method
	SecurityEnabled *bool `json:"security_enabled,omitempty"`

	// Wi-Fi密码，长度8-63个字符，包含大写字母、小写字母、数字、特殊字符中至少两种，不能和Wi-Fi名称及名称逆序相同，特殊字符只支持!~@_.?
	Password *string `json:"password,omitempty"`

	// 认证类型
	AuthenticationMethod *EquipmentWlanItemAuthenticationMethod `json:"authentication_method,omitempty"`

	// 加密方式，认证类型为WPA或者WPA2时，可选TKIP、AES
	EncrptionMethod *EquipmentWlanItemEncrptionMethod `json:"encrption_method,omitempty"`

	// 是否隐藏Wi-Fi名称
	NameHided *bool `json:"name_hided,omitempty"`
}

func (o EquipmentWlanItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentWlanItem struct{}"
	}

	return strings.Join([]string{"EquipmentWlanItem", string(data)}, " ")
}

type EquipmentWlanItemAuthenticationMethod struct {
	value string
}

type EquipmentWlanItemAuthenticationMethodEnum struct {
	WPA  EquipmentWlanItemAuthenticationMethod
	WPA2 EquipmentWlanItemAuthenticationMethod
}

func GetEquipmentWlanItemAuthenticationMethodEnum() EquipmentWlanItemAuthenticationMethodEnum {
	return EquipmentWlanItemAuthenticationMethodEnum{
		WPA: EquipmentWlanItemAuthenticationMethod{
			value: "WPA",
		},
		WPA2: EquipmentWlanItemAuthenticationMethod{
			value: "WPA2",
		},
	}
}

func (c EquipmentWlanItemAuthenticationMethod) Value() string {
	return c.value
}

func (c EquipmentWlanItemAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentWlanItemAuthenticationMethod) UnmarshalJSON(b []byte) error {
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

type EquipmentWlanItemEncrptionMethod struct {
	value string
}

type EquipmentWlanItemEncrptionMethodEnum struct {
	TKIP EquipmentWlanItemEncrptionMethod
	AES  EquipmentWlanItemEncrptionMethod
}

func GetEquipmentWlanItemEncrptionMethodEnum() EquipmentWlanItemEncrptionMethodEnum {
	return EquipmentWlanItemEncrptionMethodEnum{
		TKIP: EquipmentWlanItemEncrptionMethod{
			value: "TKIP",
		},
		AES: EquipmentWlanItemEncrptionMethod{
			value: "AES",
		},
	}
}

func (c EquipmentWlanItemEncrptionMethod) Value() string {
	return c.value
}

func (c EquipmentWlanItemEncrptionMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentWlanItemEncrptionMethod) UnmarshalJSON(b []byte) error {
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
