package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipeProcessStatus **参数解释**: 作业处理状态 - COMPLETED 已完成 - CREATING 创建中 - UPDATING 更新中 - DELETING 删除中 - UPDATING_FAILED 更新失败 - DELETING_FAILED 删除失败  **约束限制** 不涉及 **取值范围**: - COMPLETED - CREATING - UPDATING - DELETING - UPDATING_FAILED - DELETING_FAILED  **默认值** 不涉及
type PipeProcessStatus struct {
	value string
}

type PipeProcessStatusEnum struct {
	COMPLETED       PipeProcessStatus
	CREATING        PipeProcessStatus
	UPDATING        PipeProcessStatus
	DELETING        PipeProcessStatus
	CREATE_FAILED   PipeProcessStatus
	UPDATING_FAILED PipeProcessStatus
	DELETING_FAILED PipeProcessStatus
}

func GetPipeProcessStatusEnum() PipeProcessStatusEnum {
	return PipeProcessStatusEnum{
		COMPLETED: PipeProcessStatus{
			value: "COMPLETED",
		},
		CREATING: PipeProcessStatus{
			value: "CREATING",
		},
		UPDATING: PipeProcessStatus{
			value: "UPDATING",
		},
		DELETING: PipeProcessStatus{
			value: "DELETING",
		},
		CREATE_FAILED: PipeProcessStatus{
			value: "CREATE_FAILED",
		},
		UPDATING_FAILED: PipeProcessStatus{
			value: "UPDATING_FAILED",
		},
		DELETING_FAILED: PipeProcessStatus{
			value: "DELETING_FAILED",
		},
	}
}

func (c PipeProcessStatus) Value() string {
	return c.value
}

func (c PipeProcessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipeProcessStatus) UnmarshalJSON(b []byte) error {
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
