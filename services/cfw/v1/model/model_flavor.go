package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Flavor struct {

	// 防火墙版本，0：标准版，1：专业版，3：基础版，购买时，当防火墙“charge_mode”为“postPaid”时，仅支持专业版。“charge_mode”为“prePaid”时，支持标准版、专业版。
	Version *FlavorVersion `json:"version,omitempty"`

	// eip数量
	EipCount *int32 `json:"eip_count,omitempty"`

	// vpc数量
	VpcCount *int32 `json:"vpc_count,omitempty"`

	// 带宽，单位为mbps
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 日志存储，单位为byte
	LogStorage *int32 `json:"log_storage,omitempty"`

	// 默认防火墙带宽，单位为mbps，标准版为10，专业版为50，按需专业版为200
	DefaultBandwidth *int32 `json:"default_bandwidth,omitempty"`

	// 默认eip数，标准版为20，专业版为50，按需专业版为1000
	DefaultEipCount *int32 `json:"default_eip_count,omitempty"`

	// 默认日志存储，单位为byte，默认为0
	DefaultLogStorage *int32 `json:"default_log_storage,omitempty"`

	// 默认vpc数，标准版为0，专业版为2，按需专业版为5
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
