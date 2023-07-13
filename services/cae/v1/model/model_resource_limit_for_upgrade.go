package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceLimitForUpgrade struct {

	// cpu限额。
	CpuLimit *ResourceLimitForUpgradeCpuLimit `json:"cpu_limit,omitempty"`

	// 内存限额。
	MemoryLimit *ResourceLimitForUpgradeMemoryLimit `json:"memory_limit,omitempty"`
}

func (o ResourceLimitForUpgrade) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceLimitForUpgrade struct{}"
	}

	return strings.Join([]string{"ResourceLimitForUpgrade", string(data)}, " ")
}

type ResourceLimitForUpgradeCpuLimit struct {
	value string
}

type ResourceLimitForUpgradeCpuLimitEnum struct {
	E_500M  ResourceLimitForUpgradeCpuLimit
	E_1000M ResourceLimitForUpgradeCpuLimit
	E_2000M ResourceLimitForUpgradeCpuLimit
}

func GetResourceLimitForUpgradeCpuLimitEnum() ResourceLimitForUpgradeCpuLimitEnum {
	return ResourceLimitForUpgradeCpuLimitEnum{
		E_500M: ResourceLimitForUpgradeCpuLimit{
			value: "500m",
		},
		E_1000M: ResourceLimitForUpgradeCpuLimit{
			value: "1000m",
		},
		E_2000M: ResourceLimitForUpgradeCpuLimit{
			value: "2000m",
		},
	}
}

func (c ResourceLimitForUpgradeCpuLimit) Value() string {
	return c.value
}

func (c ResourceLimitForUpgradeCpuLimit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceLimitForUpgradeCpuLimit) UnmarshalJSON(b []byte) error {
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

type ResourceLimitForUpgradeMemoryLimit struct {
	value string
}

type ResourceLimitForUpgradeMemoryLimitEnum struct {
	E_1_GI ResourceLimitForUpgradeMemoryLimit
	E_2_GI ResourceLimitForUpgradeMemoryLimit
	E_4_GI ResourceLimitForUpgradeMemoryLimit
}

func GetResourceLimitForUpgradeMemoryLimitEnum() ResourceLimitForUpgradeMemoryLimitEnum {
	return ResourceLimitForUpgradeMemoryLimitEnum{
		E_1_GI: ResourceLimitForUpgradeMemoryLimit{
			value: "1Gi",
		},
		E_2_GI: ResourceLimitForUpgradeMemoryLimit{
			value: "2Gi",
		},
		E_4_GI: ResourceLimitForUpgradeMemoryLimit{
			value: "4Gi",
		},
	}
}

func (c ResourceLimitForUpgradeMemoryLimit) Value() string {
	return c.value
}

func (c ResourceLimitForUpgradeMemoryLimit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceLimitForUpgradeMemoryLimit) UnmarshalJSON(b []byte) error {
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
