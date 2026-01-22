package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateFirewallReqFlavor **参数解释**： 防火墙规格信息 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type CreateFirewallReqFlavor struct {

	// **参数解释**： 防火墙版本 **约束限制**： “charge_mode”为“prePaid”时，支持标准版、专业版。 “charge_mode”为“postPaid”时，仅支持专业版。 **取值范围**： Standard - 标准版 Professional - 专业版 **默认取值**： 不涉及
	Version CreateFirewallReqFlavorVersion `json:"version"`

	// **参数解释**： 扩展EIP数量 **约束限制**： 仅包周期场景下生效，当用户需要增加EIP防护数量时使用此参数。 **取值范围**： 0-2000 **默认取值**： 0
	ExtendEipCount *int32 `json:"extend_eip_count,omitempty"`

	// **参数解释**： 扩展带宽，步长为5 **约束限制**： 仅包周期场景下生效，当用户需要防护带宽时使用此参数。 **取值范围**： 0-5000 **默认取值**： 不涉及
	ExtendBandwidth *int32 `json:"extend_bandwidth,omitempty"`

	// **参数解释**： 扩展VPC数量 **约束限制**： 仅包周期场景下生效，当用户需要增加VPC防护数量时使用此参数。 **取值范围**： 0-100 **默认取值**： 不涉及
	ExtendVpcCount *int32 `json:"extend_vpc_count,omitempty"`
}

func (o CreateFirewallReqFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallReqFlavor struct{}"
	}

	return strings.Join([]string{"CreateFirewallReqFlavor", string(data)}, " ")
}

type CreateFirewallReqFlavorVersion struct {
	value string
}

type CreateFirewallReqFlavorVersionEnum struct {
	STANDARD     CreateFirewallReqFlavorVersion
	PROFESSIONAL CreateFirewallReqFlavorVersion
}

func GetCreateFirewallReqFlavorVersionEnum() CreateFirewallReqFlavorVersionEnum {
	return CreateFirewallReqFlavorVersionEnum{
		STANDARD: CreateFirewallReqFlavorVersion{
			value: "Standard",
		},
		PROFESSIONAL: CreateFirewallReqFlavorVersion{
			value: "Professional",
		},
	}
}

func (c CreateFirewallReqFlavorVersion) Value() string {
	return c.value
}

func (c CreateFirewallReqFlavorVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFirewallReqFlavorVersion) UnmarshalJSON(b []byte) error {
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
