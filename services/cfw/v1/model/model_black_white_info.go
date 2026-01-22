package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BlackWhiteInfo struct {

	// **参数解释**： IP地址 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	Address string `json:"address"`

	// **参数解释**： IP地址类型 **约束限制**： 不涉及  **取值范围**：  0：IPV4  1：IPV6 **默认取值**： 不涉及
	AddressType *int32 `json:"address_type,omitempty"`

	// **参数解释**： 描述 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 地址方向 **约束限制**： 不涉及  **取值范围**： 0：源地址 1：目的地址 **默认取值**： 不涉及
	Direction int32 `json:"direction"`

	// **参数解释**： 端口 **约束限制**： 不涉及  **取值范围**： 不涉及 **默认取值**： 不涉及
	Port *string `json:"port,omitempty"`

	// **参数解释**： 协议类型 **约束限制**： 不涉及  **取值范围**： -1：ANY 1：ICMP 6：TCP 17：UDP 58：ICMPV6 132：SCTP 手动类型不为空，自动类型为空 **默认取值**： 不涉及
	Protocol BlackWhiteInfoProtocol `json:"protocol"`
}

func (o BlackWhiteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteInfo struct{}"
	}

	return strings.Join([]string{"BlackWhiteInfo", string(data)}, " ")
}

type BlackWhiteInfoProtocol struct {
	value int32
}

type BlackWhiteInfoProtocolEnum struct {
	E_6       BlackWhiteInfoProtocol
	E_17      BlackWhiteInfoProtocol
	E_1       BlackWhiteInfoProtocol
	E_58      BlackWhiteInfoProtocol
	E_MINUS_1 BlackWhiteInfoProtocol
	E_132     BlackWhiteInfoProtocol
}

func GetBlackWhiteInfoProtocolEnum() BlackWhiteInfoProtocolEnum {
	return BlackWhiteInfoProtocolEnum{
		E_6: BlackWhiteInfoProtocol{
			value: 6,
		}, E_17: BlackWhiteInfoProtocol{
			value: 17,
		}, E_1: BlackWhiteInfoProtocol{
			value: 1,
		}, E_58: BlackWhiteInfoProtocol{
			value: 58,
		}, E_MINUS_1: BlackWhiteInfoProtocol{
			value: -1,
		}, E_132: BlackWhiteInfoProtocol{
			value: 132,
		},
	}
}

func (c BlackWhiteInfoProtocol) Value() int32 {
	return c.value
}

func (c BlackWhiteInfoProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BlackWhiteInfoProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
