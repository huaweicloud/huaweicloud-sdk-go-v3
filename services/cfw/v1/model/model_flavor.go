package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Flavor struct {

	// **参数解释**： 防火墙版本 **取值范围**： - 0：标准版 - 1：专业版 - 3：基础版，
	Version *FlavorVersion `json:"version,omitempty"`

	// **参数解释**： EIP数量 **取值范围**： 不涉及
	EipCount *int32 `json:"eip_count,omitempty"`

	// **参数解释**： VPC数量 **取值范围**： 不涉及
	VpcCount *int32 `json:"vpc_count,omitempty"`

	// **参数解释**： 带宽，单位为mbps **取值范围**： 不涉及
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// **参数解释**： 日志存储，单位为byte **取值范围**： 不涉及
	LogStorage *int32 `json:"log_storage,omitempty"`

	// **参数解释**： 默认防火墙带宽，单位为mbps **取值范围**： 包周期标准版为10，专业版为50，按需专业版为200
	DefaultBandwidth *int32 `json:"default_bandwidth,omitempty"`

	// **参数解释**： 默认eip数 **取值范围**： 包周期标准版为20，专业版为50，按需专业版为1000
	DefaultEipCount *int32 `json:"default_eip_count,omitempty"`

	// **参数解释**： 默认日志存储，单位为byte **取值范围**： 不涉及
	DefaultLogStorage *int32 `json:"default_log_storage,omitempty"`

	// **参数解释**： 默认vpc数 **约束限制**： 包周期标准版为0，专业版为2，按需专业版为5
	DefaultVpcCount *int32 `json:"default_vpc_count,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}

type FlavorVersion struct {
	value int32
}

type FlavorVersionEnum struct {
	E_0 FlavorVersion
	E_1 FlavorVersion
	E_2 FlavorVersion
	E_3 FlavorVersion
}

func GetFlavorVersionEnum() FlavorVersionEnum {
	return FlavorVersionEnum{
		E_0: FlavorVersion{
			value: 0,
		}, E_1: FlavorVersion{
			value: 1,
		}, E_2: FlavorVersion{
			value: 2,
		}, E_3: FlavorVersion{
			value: 3,
		},
	}
}

func (c FlavorVersion) Value() int32 {
	return c.value
}

func (c FlavorVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlavorVersion) UnmarshalJSON(b []byte) error {
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
