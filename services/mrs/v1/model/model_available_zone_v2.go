package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AvailableZoneV2 可用区信息
type AvailableZoneV2 struct {

	// 可用区编码
	Id *string `json:"id,omitempty"`

	// 可用区编码
	AzCode *string `json:"az_code,omitempty"`

	// 可用区名称
	AzName *string `json:"az_name,omitempty"`

	// 可用区id
	AzId *string `json:"az_id,omitempty"`

	// 可用区状态
	Status *string `json:"status,omitempty"`

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 可用区分组id
	AzGroupId *string `json:"az_group_id,omitempty"`

	// 当前AZ的类型 Core 核心 Satellite 卫星 Dedicated 专属 Virtual 虚拟 Edge 边缘 EdgeCental 中心边缘 Hybrid 混合云
	AzType *string `json:"az_type,omitempty"`

	AzTags *AvailableTag `json:"az_tags,omitempty"`

	// 当前可用区的类型，包括： - 0: 大云主可用区 - 21: 本地可用区 - 41: 边缘可用区
	AzCategory *int32 `json:"az_category,omitempty"`

	// 当前可用区的销售策略，包括： - charge: 计费 - notCharge: 非计费
	ChargePolicy *AvailableZoneV2ChargePolicy `json:"charge_policy,omitempty"`
}

func (o AvailableZoneV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZoneV2 struct{}"
	}

	return strings.Join([]string{"AvailableZoneV2", string(data)}, " ")
}

type AvailableZoneV2ChargePolicy struct {
	value string
}

type AvailableZoneV2ChargePolicyEnum struct {
	CHARGE     AvailableZoneV2ChargePolicy
	NOT_CHARGE AvailableZoneV2ChargePolicy
}

func GetAvailableZoneV2ChargePolicyEnum() AvailableZoneV2ChargePolicyEnum {
	return AvailableZoneV2ChargePolicyEnum{
		CHARGE: AvailableZoneV2ChargePolicy{
			value: "charge",
		},
		NOT_CHARGE: AvailableZoneV2ChargePolicy{
			value: "notCharge",
		},
	}
}

func (c AvailableZoneV2ChargePolicy) Value() string {
	return c.value
}

func (c AvailableZoneV2ChargePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AvailableZoneV2ChargePolicy) UnmarshalJSON(b []byte) error {
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
