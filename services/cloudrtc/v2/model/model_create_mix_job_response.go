package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateMixJobResponse struct {

	// 任务编号，可以用于修改、查看和停止合流任务
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 流名
	StreamName *string `json:"stream_name,omitempty" xml:"stream_name"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 房间id
	RoomId *string `json:"room_id,omitempty" xml:"room_id"`

	MixParam *MixParam `json:"mix_param,omitempty" xml:"mix_param"`

	RecordParam *RecordParam `json:"record_param,omitempty" xml:"record_param"`

	// 任务创建的时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 任务中的布局更新的时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 任务状态。  - INIT：任务正在初始化 - RUNNING：任务正在运行 - STOPPED：任务已停止
	State *CreateMixJobResponseState `json:"state,omitempty" xml:"state"`

	// 任务结束原因
	StopReason *CreateMixJobResponseStopReason `json:"stop_reason,omitempty" xml:"stop_reason"`

	// 状态描述，对state字段的一些补充说明，可用于人工查阅。
	Description *string `json:"description,omitempty" xml:"description"`

	// 任务开始时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 任务结束时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
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
