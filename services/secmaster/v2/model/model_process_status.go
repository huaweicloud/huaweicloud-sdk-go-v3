package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProcessStatus **参数解释**: 处理状态 - COMPLETED 已完成 - CREATING 创建中 - UPDATING 更新中 - ENABLING 启用中 - DISABLING 停用中 - DELETING 删除中 - CREATE_FAILED 创建失败 - UPDATE_FAILED 更新失败 - ENABLE_FAILED 启用失败 - DISABLE_FAILED 停用失败 - DELETE_FAILED 删除失败 - RECOVERING 恢复中  **约束限制** 不涉及 **取值范围**: - COMPLETED - CREATING - UPDATING - ENABLING - DISABLING - DELETING - CREATE_FAILED - UPDATE_FAILED - ENABLE_FAILED - DISABLE_FAILED - DELETE_FAILED - RECOVERING  **默认值** 不涉及
type ProcessStatus struct {
	value string
}

type ProcessStatusEnum struct {
	COMPLETED      ProcessStatus
	CREATING       ProcessStatus
	UPDATING       ProcessStatus
	ENABLING       ProcessStatus
	DISABLING      ProcessStatus
	DELETING       ProcessStatus
	CREATE_FAILED  ProcessStatus
	UPDATE_FAILED  ProcessStatus
	ENABLE_FAILED  ProcessStatus
	DISABLE_FAILED ProcessStatus
	DELETE_FAILED  ProcessStatus
	RECOVERING     ProcessStatus
}

func GetProcessStatusEnum() ProcessStatusEnum {
	return ProcessStatusEnum{
		COMPLETED: ProcessStatus{
			value: "COMPLETED",
		},
		CREATING: ProcessStatus{
			value: "CREATING",
		},
		UPDATING: ProcessStatus{
			value: "UPDATING",
		},
		ENABLING: ProcessStatus{
			value: "ENABLING",
		},
		DISABLING: ProcessStatus{
			value: "DISABLING",
		},
		DELETING: ProcessStatus{
			value: "DELETING",
		},
		CREATE_FAILED: ProcessStatus{
			value: "CREATE_FAILED",
		},
		UPDATE_FAILED: ProcessStatus{
			value: "UPDATE_FAILED",
		},
		ENABLE_FAILED: ProcessStatus{
			value: "ENABLE_FAILED",
		},
		DISABLE_FAILED: ProcessStatus{
			value: "DISABLE_FAILED",
		},
		DELETE_FAILED: ProcessStatus{
			value: "DELETE_FAILED",
		},
		RECOVERING: ProcessStatus{
			value: "RECOVERING",
		},
	}
}

func (c ProcessStatus) Value() string {
	return c.value
}

func (c ProcessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProcessStatus) UnmarshalJSON(b []byte) error {
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
