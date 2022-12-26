package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateWorkflowTriggerStatusRequest struct {

	// 任务id，待修改任务的id。
	WorkflowId string `json:"workflow_id"`

	// 启动或暂停任务的定时执行。enable为开启定时任务，disable为关闭定时任务
	Action UpdateWorkflowTriggerStatusRequestAction `json:"action"`
}

func (o UpdateWorkflowTriggerStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowTriggerStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowTriggerStatusRequest", string(data)}, " ")
}

type UpdateWorkflowTriggerStatusRequestAction struct {
	value string
}

type UpdateWorkflowTriggerStatusRequestActionEnum struct {
	ENABLE  UpdateWorkflowTriggerStatusRequestAction
	DISABLE UpdateWorkflowTriggerStatusRequestAction
}

func GetUpdateWorkflowTriggerStatusRequestActionEnum() UpdateWorkflowTriggerStatusRequestActionEnum {
	return UpdateWorkflowTriggerStatusRequestActionEnum{
		ENABLE: UpdateWorkflowTriggerStatusRequestAction{
			value: "enable",
		},
		DISABLE: UpdateWorkflowTriggerStatusRequestAction{
			value: "disable",
		},
	}
}

func (c UpdateWorkflowTriggerStatusRequestAction) Value() string {
	return c.value
}

func (c UpdateWorkflowTriggerStatusRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWorkflowTriggerStatusRequestAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
