package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowMixJobResponse struct {

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
	State *ShowMixJobResponseState `json:"state,omitempty" xml:"state"`

	// 任务结束原因
	StopReason *ShowMixJobResponseStopReason `json:"stop_reason,omitempty" xml:"stop_reason"`

	// 状态描述，对state字段的一些补充说明，可用于人工查阅。
	Description *string `json:"description,omitempty" xml:"description"`

	// 任务开始时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 任务结束时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMixJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMixJobResponse struct{}"
	}

	return strings.Join([]string{"ShowMixJobResponse", string(data)}, " ")
}

type ShowMixJobResponseState struct {
	value string
}

type ShowMixJobResponseStateEnum struct {
	INIT    ShowMixJobResponseState
	RUNNING ShowMixJobResponseState
	STOPPED ShowMixJobResponseState
}

func GetShowMixJobResponseStateEnum() ShowMixJobResponseStateEnum {
	return ShowMixJobResponseStateEnum{
		INIT: ShowMixJobResponseState{
			value: "INIT",
		},
		RUNNING: ShowMixJobResponseState{
			value: "RUNNING",
		},
		STOPPED: ShowMixJobResponseState{
			value: "STOPPED",
		},
	}
}

func (c ShowMixJobResponseState) Value() string {
	return c.value
}

func (c ShowMixJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMixJobResponseState) UnmarshalJSON(b []byte) error {
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

type ShowMixJobResponseStopReason struct {
	value string
}

type ShowMixJobResponseStopReasonEnum struct {
	TENANT_STOP          ShowMixJobResponseStopReason
	EXCEED_MAX_IDLE_TIME ShowMixJobResponseStopReason
	INTERNAL_ERROR       ShowMixJobResponseStopReason
}

func GetShowMixJobResponseStopReasonEnum() ShowMixJobResponseStopReasonEnum {
	return ShowMixJobResponseStopReasonEnum{
		TENANT_STOP: ShowMixJobResponseStopReason{
			value: "TENANT_STOP",
		},
		EXCEED_MAX_IDLE_TIME: ShowMixJobResponseStopReason{
			value: "EXCEED_MAX_IDLE_TIME",
		},
		INTERNAL_ERROR: ShowMixJobResponseStopReason{
			value: "INTERNAL_ERROR",
		},
	}
}

func (c ShowMixJobResponseStopReason) Value() string {
	return c.value
}

func (c ShowMixJobResponseStopReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMixJobResponseStopReason) UnmarshalJSON(b []byte) error {
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
