package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstanceState **参数解释**： 实例状态 **约束限制**： 不涉及 **取值范围**： - pending：实例正在启动（分配资源/启动操作系统） - running：实例正常运行（可接受SSH/RDP连接） - stopping：实例正在关闭（停止/休眠过渡状态） - stopped： 实例已完全关闭（存储卷保留） - reinstalling：实例正在重装中 - shutting-down：实例正在终止（删除流程中） - terminated：实例已终止（资源完全删除，不可恢复） - failed：实例处于失败状态，对于reinstall操作可重试，其它操作不可重试并清除相关资源 - modifyIping：实例正在修改ip中 - switch-installing：实例正在切换系统中  ```mermaid stateDiagram-v2      [*] --> pending : RunInstances/CreateInstance     pending --> running : Provision Finished     reinstalling --> running : Provision Finished     running --> stopping : PowerOff/PowerReboot     stopping --> stopped : PowerOff Finished     stopped --> running : PowerOn | ModifyIP     stopped --> stopped : ChangePassword     pending --> shutting_down : Abort by DeleteInstance     running --> shutting_down : DeleteInstance     running --> switch-installing : SwitchOS     switch-installing --> running : SwitchOS finished     switch-installing --> failed : SwitchOS failed     stopped --> shutting_down : DeleteInstance     stopped --> reinstalling : ReinstallOS     shutting_down -->terminated: DeleteInstance Finished     pending --> failed : Error     shutting_down --> failed : Error     reinstalling --> failed : Error     failed --> reinstalling: Retry     running --> modifyIping : Modify Ip     modifyIping --> running : Modify Ip finished     modifyIping --> failed : Modify Ip failed     failed --> modifyIping : Retrey ```    **默认取值**：   不涉及
type InstanceState struct {
	value string
}

type InstanceStateEnum struct {
	PENDING           InstanceState
	RUNNING           InstanceState
	STOPPING          InstanceState
	STOPPED           InstanceState
	REINSTALLING      InstanceState
	SHUTTING_DOWN     InstanceState
	TERMINATED        InstanceState
	FAILED            InstanceState
	MODIFY_IPING      InstanceState
	SWITCH_INSTALLING InstanceState
}

func GetInstanceStateEnum() InstanceStateEnum {
	return InstanceStateEnum{
		PENDING: InstanceState{
			value: "pending",
		},
		RUNNING: InstanceState{
			value: "running",
		},
		STOPPING: InstanceState{
			value: "stopping",
		},
		STOPPED: InstanceState{
			value: "stopped",
		},
		REINSTALLING: InstanceState{
			value: "reinstalling",
		},
		SHUTTING_DOWN: InstanceState{
			value: "shutting-down",
		},
		TERMINATED: InstanceState{
			value: "terminated",
		},
		FAILED: InstanceState{
			value: "failed",
		},
		MODIFY_IPING: InstanceState{
			value: "modifyIping",
		},
		SWITCH_INSTALLING: InstanceState{
			value: "switch-installing",
		},
	}
}

func (c InstanceState) Value() string {
	return c.value
}

func (c InstanceState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceState) UnmarshalJSON(b []byte) error {
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
