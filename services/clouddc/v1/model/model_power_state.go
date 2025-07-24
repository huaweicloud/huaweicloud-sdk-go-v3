package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PowerState **参数解释**： 电源状态 power_state 会根据不同的操作和事件发生转换，常见的状态转换流程如下：   - 开机流程：off -> powering-on -> on   - 关机流程：on -> powering-off -> off   - 重启流程：on -> rebooting -> on  **约束限制**： 不涉及 **取值范围**： - on：表示节点的电源已开启，硬件处于通电状态，操作系统正在运行或者可以正常启动。这意味着节点能够执行计算任务，为上层应用提供服务。 示例场景：当用户在 Ironic 中创建并激活一个节点，或者手动开启节点电源后，节点的 power_state 会变为 power on。 - off：表明节点的电源已关闭，硬件停止供电，所有组件处于非工作状态，无法执行任何计算任务。 示例场景：在维护节点或者不需要使用节点资源时，管理员可以将节点的电源关闭，此时 power_state 变为 power off。 - rebooting：节点正在进行重启操作，即先关闭电源，然后再重新开启。在这个过程中，节点会经历硬件初始化和操作系统启动等步骤。   示例场景：当管理员通过 Ironic API 发送重启节点的指令后，节点的 power_state 会暂时变为 rebooting，直到重启完成。 - powering-on：节点正在开启电源的过程中，此时硬件开始通电，但操作系统可能还未完全启动。 示例场景：当管理员发送开机指令后，节点会进入 powering on 状态，直到操作系统成功启动，power_state 变为 power on。 - powering-off：节点正在关闭电源的过程中，操作系统会进行一些清理工作，如保存数据、关闭服务等，然后切断硬件的电源供应。  示例场景：当管理员发送关机指令后，节点会进入 powering off 状态，直到电源完全关闭，power_state 变为  off。  **默认取值**： 不涉及
type PowerState struct {
	value string
}

type PowerStateEnum struct {
	ON           PowerState
	POWERING_ON  PowerState
	OFF          PowerState
	REBOOTING    PowerState
	POWERING_OFF PowerState
}

func GetPowerStateEnum() PowerStateEnum {
	return PowerStateEnum{
		ON: PowerState{
			value: "on",
		},
		POWERING_ON: PowerState{
			value: "powering-on",
		},
		OFF: PowerState{
			value: "off",
		},
		REBOOTING: PowerState{
			value: "rebooting",
		},
		POWERING_OFF: PowerState{
			value: "powering-off",
		},
	}
}

func (c PowerState) Value() string {
	return c.value
}

func (c PowerState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PowerState) UnmarshalJSON(b []byte) error {
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
