package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstanceOperationStatus **参数解释**:  操作状态 **约束限制**:  不涉及 **取值范围**:  - install-processing: 安装OS中 - install-succeed: 安装OS成功 - install-failed: 安装OS失败 - reinstall-processing: 重装OS中 - reinstall-succeed: 重装OS成功 - reinstall-failed: 重装OS失败 - switch-install-processing: 切换OS中 - switch-install-succeed: 切换OS成功 - switch-install-failed: 切换OS失败 - modify-ip-processing: 修改IP地址中 - modify-ip-succeed: 修改IP地址成功 - modify-ip-failed: 修改IP地址失败 - uninstall-processing: 卸载OS中 - uninstall-succeed: 卸载OS成功 - uninstall-failed: 卸载OS失败  **默认取值**:  不涉及
type InstanceOperationStatus struct {
	value string
}

type InstanceOperationStatusEnum struct {
	INSTALL_PROCESSING        InstanceOperationStatus
	INSTALL_SUCCEED           InstanceOperationStatus
	INSTALL_FAILED            InstanceOperationStatus
	REINSTALL_PROCESSING      InstanceOperationStatus
	REINSTALL_SUCCEED         InstanceOperationStatus
	REINSTALL_FAILED          InstanceOperationStatus
	SWITCH_INSTALL_PROCESSING InstanceOperationStatus
	SWITCH_INSTALL_SUCCEED    InstanceOperationStatus
	SWITCH_INSTALL_FAILED     InstanceOperationStatus
	MODIFY_IP_PROCESSING      InstanceOperationStatus
	MODIFY_IP_SUCCEED         InstanceOperationStatus
	MODIFY_IP_FAILED          InstanceOperationStatus
	UNINSTALL_PROCESSING      InstanceOperationStatus
	UNINSTALL_SUCCEED         InstanceOperationStatus
	UNINSTALL_FAILED          InstanceOperationStatus
}

func GetInstanceOperationStatusEnum() InstanceOperationStatusEnum {
	return InstanceOperationStatusEnum{
		INSTALL_PROCESSING: InstanceOperationStatus{
			value: "install-processing",
		},
		INSTALL_SUCCEED: InstanceOperationStatus{
			value: "install-succeed",
		},
		INSTALL_FAILED: InstanceOperationStatus{
			value: "install-failed",
		},
		REINSTALL_PROCESSING: InstanceOperationStatus{
			value: "reinstall-processing",
		},
		REINSTALL_SUCCEED: InstanceOperationStatus{
			value: "reinstall-succeed",
		},
		REINSTALL_FAILED: InstanceOperationStatus{
			value: "reinstall-failed",
		},
		SWITCH_INSTALL_PROCESSING: InstanceOperationStatus{
			value: "switch-install-processing",
		},
		SWITCH_INSTALL_SUCCEED: InstanceOperationStatus{
			value: "switch-install-succeed",
		},
		SWITCH_INSTALL_FAILED: InstanceOperationStatus{
			value: "switch-install-failed",
		},
		MODIFY_IP_PROCESSING: InstanceOperationStatus{
			value: "modify-ip-processing",
		},
		MODIFY_IP_SUCCEED: InstanceOperationStatus{
			value: "modify-ip-succeed",
		},
		MODIFY_IP_FAILED: InstanceOperationStatus{
			value: "modify-ip-failed",
		},
		UNINSTALL_PROCESSING: InstanceOperationStatus{
			value: "uninstall-processing",
		},
		UNINSTALL_SUCCEED: InstanceOperationStatus{
			value: "uninstall-succeed",
		},
		UNINSTALL_FAILED: InstanceOperationStatus{
			value: "uninstall-failed",
		},
	}
}

func (c InstanceOperationStatus) Value() string {
	return c.value
}

func (c InstanceOperationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceOperationStatus) UnmarshalJSON(b []byte) error {
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
