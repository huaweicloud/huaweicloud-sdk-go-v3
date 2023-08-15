package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Flag 特殊标识，用于前端使用
type Flag struct {

	// 是否开启pci_3ds合规认证   - true：开启   - false：不开启
	Pci3ds *FlagPci3ds `json:"pci_3ds,omitempty"`

	// 是否开启pci_dss合规认证   - true：开启   - false：不开启
	PciDss *FlagPciDss `json:"pci_dss,omitempty"`

	// old：代表域名使用的老的cname，new：代表域名使用新的cname
	Cname *FlagCname `json:"cname,omitempty"`

	// 域名是否开启ipv6   - true：支持   - false：不支持
	IsDualAz *FlagIsDualAz `json:"is_dual_az,omitempty"`

	// 域名是否开启ipv6   - true：支持   - false：不支持
	Ipv6 *FlagIpv6 `json:"ipv6,omitempty"`
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

func (c FlagPci3ds) Value() string {
	return c.value
}

func (c FlagPci3ds) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagPci3ds) UnmarshalJSON(b []byte) error {
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

func (c FlagPciDss) Value() string {
	return c.value
}

func (c FlagPciDss) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagPciDss) UnmarshalJSON(b []byte) error {
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

func (c FlagCname) Value() string {
	return c.value
}

func (c FlagCname) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagCname) UnmarshalJSON(b []byte) error {
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

func (c FlagIsDualAz) Value() string {
	return c.value
}

func (c FlagIsDualAz) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagIsDualAz) UnmarshalJSON(b []byte) error {
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

type FlagIpv6 struct {
	value string
}

type FlagIpv6Enum struct {
	TRUE  FlagIpv6
	FALSE FlagIpv6
}

func GetFlagIpv6Enum() FlagIpv6Enum {
	return FlagIpv6Enum{
		TRUE: FlagIpv6{
			value: "true",
		},
		FALSE: FlagIpv6{
			value: "false",
		},
	}
}

func (c FlagIpv6) Value() string {
	return c.value
}

func (c FlagIpv6) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlagIpv6) UnmarshalJSON(b []byte) error {
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
