package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WAF支持的认证项
type Flag struct {
	// true:通过pci_3ds标准认证,false:未通过pci_3ds标准认证

	Pci3ds *FlagPci3ds `json:"pci_3ds,omitempty"`
	// true:通过pci_dss标准认证,false:未通过pci_dss标准认证

	PciDss *FlagPciDss `json:"pci_dss,omitempty"`
	// old：代表域名使用的老的cname，new：代表域名使用新的cname

	Cname *FlagCname `json:"cname,omitempty"`
	// true：WAF支持多可用区容灾,false：WAF不支持多可用区容灾

	IsDualAz *FlagIsDualAz `json:"is_dual_az,omitempty"`
}

func (o Flag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flag struct{}"
	}

	return strings.Join([]string{"Flag", string(data)}, " ")
}

type FlagPci3ds struct {
	value string
}

type FlagPci3dsEnum struct {
	TRUE  FlagPci3ds
	FALSE FlagPci3ds
}

func GetFlagPci3dsEnum() FlagPci3dsEnum {
	return FlagPci3dsEnum{
		TRUE: FlagPci3ds{
			value: "true",
		},
		FALSE: FlagPci3ds{
			value: "false",
		},
	}
}

func (c FlagPci3ds) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagPci3ds) UnmarshalJSON(b []byte) error {
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

type FlagPciDss struct {
	value string
}

type FlagPciDssEnum struct {
	TRUE  FlagPciDss
	FALSE FlagPciDss
}

func GetFlagPciDssEnum() FlagPciDssEnum {
	return FlagPciDssEnum{
		TRUE: FlagPciDss{
			value: "true",
		},
		FALSE: FlagPciDss{
			value: "false",
		},
	}
}

func (c FlagPciDss) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagPciDss) UnmarshalJSON(b []byte) error {
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

type FlagCname struct {
	value string
}

type FlagCnameEnum struct {
	OLD FlagCname
	NEW FlagCname
}

func GetFlagCnameEnum() FlagCnameEnum {
	return FlagCnameEnum{
		OLD: FlagCname{
			value: "old",
		},
		NEW: FlagCname{
			value: "new",
		},
	}
}

func (c FlagCname) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagCname) UnmarshalJSON(b []byte) error {
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

type FlagIsDualAz struct {
	value string
}

type FlagIsDualAzEnum struct {
	TRUE  FlagIsDualAz
	FALSE FlagIsDualAz
}

func GetFlagIsDualAzEnum() FlagIsDualAzEnum {
	return FlagIsDualAzEnum{
		TRUE: FlagIsDualAz{
			value: "true",
		},
		FALSE: FlagIsDualAz{
			value: "false",
		},
	}
}

func (c FlagIsDualAz) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagIsDualAz) UnmarshalJSON(b []byte) error {
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
