package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PhysicalServerOperationStatus **约束限制**： 不涉及 **取值范围**： - power-on-processing: 节点正在开启电源的过程中，此时硬件开始通电，但操作系统可能还未完全启动。 - power-on-succeed: 开启电源成功。 - power-on-failed: 开启电源失败。 - power-off-processing: 节点正在关闭电源的过程中，操作系统会进行一些清理工作，如保存数据、关闭服务等，然后切断硬件的电源供应。 - power-off-succeed: 关闭电源成功。 - power-off-failed: 关闭电源失败。 - reboot-processing: 节点正在进行重启操作，即先关闭电源，然后再重新开启。在这个过程中，节点会经历硬件初始化和操作系统启动等步骤。 - reboot-succeed: 重启操作成功。 - reboot-failed: 重启操作失败。  **默认取值**： 不涉及
type PhysicalServerOperationStatus struct {
	value string
}

type PhysicalServerOperationStatusEnum struct {
	POWER_ON_PROCESSING  PhysicalServerOperationStatus
	POWER_ON_SUCCEED     PhysicalServerOperationStatus
	POWER_ON_FAILED      PhysicalServerOperationStatus
	POWER_OFF_PROCESSING PhysicalServerOperationStatus
	POWER_OFF_SUCCEED    PhysicalServerOperationStatus
	POWER_OFF_FAILED     PhysicalServerOperationStatus
	REBOOT_PROCESSING    PhysicalServerOperationStatus
	REBOOT_SUCCEED       PhysicalServerOperationStatus
	REBOOT_FAILED        PhysicalServerOperationStatus
}

func GetPhysicalServerOperationStatusEnum() PhysicalServerOperationStatusEnum {
	return PhysicalServerOperationStatusEnum{
		POWER_ON_PROCESSING: PhysicalServerOperationStatus{
			value: "power-on-processing",
		},
		POWER_ON_SUCCEED: PhysicalServerOperationStatus{
			value: "power-on-succeed",
		},
		POWER_ON_FAILED: PhysicalServerOperationStatus{
			value: "power-on-failed",
		},
		POWER_OFF_PROCESSING: PhysicalServerOperationStatus{
			value: "power-off-processing",
		},
		POWER_OFF_SUCCEED: PhysicalServerOperationStatus{
			value: "power-off-succeed",
		},
		POWER_OFF_FAILED: PhysicalServerOperationStatus{
			value: "power-off-failed",
		},
		REBOOT_PROCESSING: PhysicalServerOperationStatus{
			value: "reboot-processing",
		},
		REBOOT_SUCCEED: PhysicalServerOperationStatus{
			value: "reboot-succeed",
		},
		REBOOT_FAILED: PhysicalServerOperationStatus{
			value: "reboot-failed",
		},
	}
}

func (c PhysicalServerOperationStatus) Value() string {
	return c.value
}

func (c PhysicalServerOperationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PhysicalServerOperationStatus) UnmarshalJSON(b []byte) error {
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
