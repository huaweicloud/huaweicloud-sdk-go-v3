package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobProcessStatus **参数解释**: 作业处理状态 - COMPLETED 已完成 - CREATING 创建中 - UPDATING 更新中 - ENABLING 启用中 - DISABLING 停用中 - DELETING 删除中 - CREATE_FAILED 创建失败 - UPDATE_FAILED 更新失败 - ENABLE_FAILED 启用失败 - DISABLE_FAILED 停用失败 - DELETE_FAILED 删除失败 - RECOVERING 恢复中  **约束限制** 不涉及 **取值范围**: - COMPLETED - CREATING - UPDATING - ENABLING - DISABLING - DELETING - CREATE_FAILED - UPDATE_FAILED - ENABLE_FAILED - DISABLE_FAILED - DELETE_FAILED - RECOVERING  **默认值** 不涉及
type JobProcessStatus struct {
	value string
}

type JobProcessStatusEnum struct {
	COMPLETED      JobProcessStatus
	CREATING       JobProcessStatus
	UPDATING       JobProcessStatus
	ENABLING       JobProcessStatus
	DISABLING      JobProcessStatus
	DELETING       JobProcessStatus
	CREATE_FAILED  JobProcessStatus
	UPDATE_FAILED  JobProcessStatus
	ENABLE_FAILED  JobProcessStatus
	DISABLE_FAILED JobProcessStatus
	DELETE_FAILED  JobProcessStatus
	RECOVERING     JobProcessStatus
}

func GetJobProcessStatusEnum() JobProcessStatusEnum {
	return JobProcessStatusEnum{
		COMPLETED: JobProcessStatus{
			value: "COMPLETED",
		},
		CREATING: JobProcessStatus{
			value: "CREATING",
		},
		UPDATING: JobProcessStatus{
			value: "UPDATING",
		},
		ENABLING: JobProcessStatus{
			value: "ENABLING",
		},
		DISABLING: JobProcessStatus{
			value: "DISABLING",
		},
		DELETING: JobProcessStatus{
			value: "DELETING",
		},
		CREATE_FAILED: JobProcessStatus{
			value: "CREATE_FAILED",
		},
		UPDATE_FAILED: JobProcessStatus{
			value: "UPDATE_FAILED",
		},
		ENABLE_FAILED: JobProcessStatus{
			value: "ENABLE_FAILED",
		},
		DISABLE_FAILED: JobProcessStatus{
			value: "DISABLE_FAILED",
		},
		DELETE_FAILED: JobProcessStatus{
			value: "DELETE_FAILED",
		},
		RECOVERING: JobProcessStatus{
			value: "RECOVERING",
		},
	}
}

func (c JobProcessStatus) Value() string {
	return c.value
}

func (c JobProcessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobProcessStatus) UnmarshalJSON(b []byte) error {
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
