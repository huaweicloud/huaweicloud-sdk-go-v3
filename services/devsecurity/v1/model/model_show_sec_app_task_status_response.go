package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecAppTaskStatusResponse Response Object
type ShowSecAppTaskStatusResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 创建任务的时间
	CreateTime *string `json:"create_time,omitempty"`

	// 任务启动的时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束的时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态:   * WAITING - 等待   * RUNNING - 进行   * SUCCESS - 完成   * FAILURE - 失败   * STOP - 停止   * DELETED - 删除
	TaskStatus *ShowSecAppTaskStatusResponseTaskStatus `json:"task_status,omitempty"`

	// 任务失败时返回失败原因
	FailReason     *string `json:"fail_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecAppTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecAppTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSecAppTaskStatusResponse", string(data)}, " ")
}

type ShowSecAppTaskStatusResponseTaskStatus struct {
	value string
}

type ShowSecAppTaskStatusResponseTaskStatusEnum struct {
	WAITING ShowSecAppTaskStatusResponseTaskStatus
	RUNNING ShowSecAppTaskStatusResponseTaskStatus
	SUCCESS ShowSecAppTaskStatusResponseTaskStatus
	FAILURE ShowSecAppTaskStatusResponseTaskStatus
	STOP    ShowSecAppTaskStatusResponseTaskStatus
	DELETED ShowSecAppTaskStatusResponseTaskStatus
}

func GetShowSecAppTaskStatusResponseTaskStatusEnum() ShowSecAppTaskStatusResponseTaskStatusEnum {
	return ShowSecAppTaskStatusResponseTaskStatusEnum{
		WAITING: ShowSecAppTaskStatusResponseTaskStatus{
			value: "WAITING",
		},
		RUNNING: ShowSecAppTaskStatusResponseTaskStatus{
			value: "RUNNING",
		},
		SUCCESS: ShowSecAppTaskStatusResponseTaskStatus{
			value: "SUCCESS",
		},
		FAILURE: ShowSecAppTaskStatusResponseTaskStatus{
			value: "FAILURE",
		},
		STOP: ShowSecAppTaskStatusResponseTaskStatus{
			value: "STOP",
		},
		DELETED: ShowSecAppTaskStatusResponseTaskStatus{
			value: "DELETED",
		},
	}
}

func (c ShowSecAppTaskStatusResponseTaskStatus) Value() string {
	return c.value
}

func (c ShowSecAppTaskStatusResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecAppTaskStatusResponseTaskStatus) UnmarshalJSON(b []byte) error {
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
