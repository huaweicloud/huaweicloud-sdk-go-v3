package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateMixJobResponse struct {

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
	State *UpdateMixJobResponseState `json:"state,omitempty"`

	// 任务结束原因 - TENANT_STOP - EXCEED_MAX_IDLE_TIME - INTERNAL_ERROR
	StopReason *UpdateMixJobResponseStopReason `json:"stop_reason,omitempty"`

	// 状态描述，对state字段的一些补充说明，可用于人工查阅。
	Description *string `json:"description,omitempty"`

	// 任务开始时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StopTime *string `json:"stop_time,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMixJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMixJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateMixJobResponse", string(data)}, " ")
}

type UpdateMixJobResponseState struct {
	value string
}

type UpdateMixJobResponseStateEnum struct {
	INIT    UpdateMixJobResponseState
	RUNNING UpdateMixJobResponseState
	STOPPED UpdateMixJobResponseState
}

func GetUpdateMixJobResponseStateEnum() UpdateMixJobResponseStateEnum {
	return UpdateMixJobResponseStateEnum{
		INIT: UpdateMixJobResponseState{
			value: "INIT",
		},
		RUNNING: UpdateMixJobResponseState{
			value: "RUNNING",
		},
		STOPPED: UpdateMixJobResponseState{
			value: "STOPPED",
		},
	}
}

func (c UpdateMixJobResponseState) Value() string {
	return c.value
}

func (c UpdateMixJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMixJobResponseState) UnmarshalJSON(b []byte) error {
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

type UpdateMixJobResponseStopReason struct {
	value string
}

type UpdateMixJobResponseStopReasonEnum struct {
	TENANT_STOP          UpdateMixJobResponseStopReason
	EXCEED_MAX_IDLE_TIME UpdateMixJobResponseStopReason
	INTERNAL_ERROR       UpdateMixJobResponseStopReason
}

func GetUpdateMixJobResponseStopReasonEnum() UpdateMixJobResponseStopReasonEnum {
	return UpdateMixJobResponseStopReasonEnum{
		TENANT_STOP: UpdateMixJobResponseStopReason{
			value: "TENANT_STOP",
		},
		EXCEED_MAX_IDLE_TIME: UpdateMixJobResponseStopReason{
			value: "EXCEED_MAX_IDLE_TIME",
		},
		INTERNAL_ERROR: UpdateMixJobResponseStopReason{
			value: "INTERNAL_ERROR",
		},
	}
}

func (c UpdateMixJobResponseStopReason) Value() string {
	return c.value
}

func (c UpdateMixJobResponseStopReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMixJobResponseStopReason) UnmarshalJSON(b []byte) error {
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
