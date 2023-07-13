package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMixJobResponse Response Object
type CreateMixJobResponse struct {

	// 任务编号，可以用于修改、查看和停止合流任务
	JobId *string `json:"job_id,omitempty"`

	// 流名
	StreamName *string `json:"stream_name,omitempty"`

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 房间id
	RoomId *string `json:"room_id,omitempty"`

	MixParam *MixParam `json:"mix_param,omitempty"`

	RecordParam *RecordParam `json:"record_param,omitempty"`

	// 任务创建的时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty"`

	// 任务中的布局更新的时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty"`

	// 任务状态。  - INIT：任务正在初始化 - RUNNING：任务正在运行 - STOPPED：任务已停止
	State *CreateMixJobResponseState `json:"state,omitempty"`

	// 任务结束原因 - TENANT_STOP - EXCEED_MAX_IDLE_TIME - INTERNAL_ERROR
	StopReason *CreateMixJobResponseStopReason `json:"stop_reason,omitempty"`

	// 状态描述，对state字段的一些补充说明，可用于人工查阅。
	Description *string `json:"description,omitempty"`

	// 任务开始时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StopTime *string `json:"stop_time,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMixJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMixJobResponse struct{}"
	}

	return strings.Join([]string{"CreateMixJobResponse", string(data)}, " ")
}

type CreateMixJobResponseState struct {
	value string
}

type CreateMixJobResponseStateEnum struct {
	INIT    CreateMixJobResponseState
	RUNNING CreateMixJobResponseState
	STOPPED CreateMixJobResponseState
}

func GetCreateMixJobResponseStateEnum() CreateMixJobResponseStateEnum {
	return CreateMixJobResponseStateEnum{
		INIT: CreateMixJobResponseState{
			value: "INIT",
		},
		RUNNING: CreateMixJobResponseState{
			value: "RUNNING",
		},
		STOPPED: CreateMixJobResponseState{
			value: "STOPPED",
		},
	}
}

func (c CreateMixJobResponseState) Value() string {
	return c.value
}

func (c CreateMixJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMixJobResponseState) UnmarshalJSON(b []byte) error {
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

type CreateMixJobResponseStopReason struct {
	value string
}

type CreateMixJobResponseStopReasonEnum struct {
	TENANT_STOP          CreateMixJobResponseStopReason
	EXCEED_MAX_IDLE_TIME CreateMixJobResponseStopReason
	INTERNAL_ERROR       CreateMixJobResponseStopReason
}

func GetCreateMixJobResponseStopReasonEnum() CreateMixJobResponseStopReasonEnum {
	return CreateMixJobResponseStopReasonEnum{
		TENANT_STOP: CreateMixJobResponseStopReason{
			value: "TENANT_STOP",
		},
		EXCEED_MAX_IDLE_TIME: CreateMixJobResponseStopReason{
			value: "EXCEED_MAX_IDLE_TIME",
		},
		INTERNAL_ERROR: CreateMixJobResponseStopReason{
			value: "INTERNAL_ERROR",
		},
	}
}

func (c CreateMixJobResponseStopReason) Value() string {
	return c.value
}

func (c CreateMixJobResponseStopReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMixJobResponseStopReason) UnmarshalJSON(b []byte) error {
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
