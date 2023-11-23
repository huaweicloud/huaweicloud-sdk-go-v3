package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Flavor struct {

	// 防火墙版本，0：标准版，1：专业版，2：铂金版，3：基础版，购买时，当防火墙“charge_mode”为“postPaid”时，仅支持专业版。“charge_mode”为“prePaid”时，支持标准版、专业版。
	Version *FlavorVersion `json:"version,omitempty"`

	// eip数量
	EipCount *int32 `json:"eip_count,omitempty"`

	// vpc数量
	VpcCount *int32 `json:"vpc_count,omitempty"`

	// 带宽
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 日志存储
	LogStorage *int32 `json:"log_storage,omitempty"`

	// 目前的会话数
	SessionConcurrent *int32 `json:"session_concurrent,omitempty"`

	// 创建会话数
	SessionCreate *int32 `json:"session_create,omitempty"`

	// 总计规则数
	TotalRuleCount *int32 `json:"total_rule_count,omitempty"`

	// 已使用规则数
	UsedRuleCount *int32 `json:"used_rule_count,omitempty"`

	// vpc间带宽
	VpcBandwith *int32 `json:"vpc_bandwith,omitempty"`
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
