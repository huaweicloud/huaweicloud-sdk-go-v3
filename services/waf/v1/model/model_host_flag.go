package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HostFlag struct {

	// 是否开启pci_3ds合规认证，该参数需要与tls和cipher参数同时使用，且tls参数值需要设置为TLS v1.2，cipher参数值设置为cipher_2。注：pci_3ds合规认证开启后不支持关闭，在开启pci_3ds合规认证前，请先阅读帮助中心Web应用防火墙WAF文档中关于pci_3ds合规认证的说明   - true：开启   - false：不开启
	Pci3ds *HostFlagPci3ds `json:"pci_3ds,omitempty"`

	// 是否开启pci_dss合规认证，该参数需要与tls和cipher参数同时使用，且tls参数值需要设置为TLS v1.2，cipher参数值设置为cipher_2。注：在开启pci_dss合规认证前，请先阅读帮助中心Web应用防火墙WAF文档中关于pci_dss合规认证的说明   - true：开启   - false：不开启
	PciDss *HostFlagPciDss `json:"pci_dss,omitempty"`
}

func (o HostFlag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostFlag struct{}"
	}

	return strings.Join([]string{"HostFlag", string(data)}, " ")
}

type HostFlagPci3ds struct {
	value string
}

type HostFlagPci3dsEnum struct {
	TRUE  HostFlagPci3ds
	FALSE HostFlagPci3ds
}

func GetHostFlagPci3dsEnum() HostFlagPci3dsEnum {
	return HostFlagPci3dsEnum{
		TRUE: HostFlagPci3ds{
			value: "true",
		},
		FALSE: HostFlagPci3ds{
			value: "false",
		},
	}
}

func (c HostFlagPci3ds) Value() string {
	return c.value
}

func (c HostFlagPci3ds) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostFlagPci3ds) UnmarshalJSON(b []byte) error {
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

type HostFlagPciDss struct {
	value string
}

type HostFlagPciDssEnum struct {
	TRUE  HostFlagPciDss
	FALSE HostFlagPciDss
}

func GetHostFlagPciDssEnum() HostFlagPciDssEnum {
	return HostFlagPciDssEnum{
		TRUE: HostFlagPciDss{
			value: "true",
		},
		FALSE: HostFlagPciDss{
			value: "false",
		},
	}
}

func (c HostFlagPciDss) Value() string {
	return c.value
}

func (c HostFlagPciDss) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HostFlagPciDss) UnmarshalJSON(b []byte) error {
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
