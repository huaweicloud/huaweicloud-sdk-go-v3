package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateIndividualStreamJobResponse Response Object
type UpdateIndividualStreamJobResponse struct {

	// 任务编号
	JobId *string `json:"job_id,omitempty"`

	// 流名
	StreamName *string `json:"stream_name,omitempty"`

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 房间id
	RoomId *string `json:"room_id,omitempty"`

	// 选看的用户id，单个录制任务内保证唯一
	UserId *string `json:"user_id,omitempty"`

	//  是否录制音频。  - true：录制音频 - false：不录制音频  缺省为true。
	IsRecordAudio *bool `json:"is_record_audio,omitempty"`

	// 标识视频流的类型，可选摄像头流或者屏幕分享流，未填写表示不录制视频。  - CAMERASTREAM：摄像头视频流 - SCREENSTREAM：屏幕分享视频流  默认为CAMERASTREAM。
	VideoType *UpdateIndividualStreamJobResponseVideoType `json:"video_type,omitempty"`

	// 指定窗口拉取的分辨率档位。  - LD - SD - HD - FHD  缺省为FHD。
	SelectStreamType *UpdateIndividualStreamJobResponseSelectStreamType `json:"select_stream_type,omitempty"`

	// 最长空闲频道时间。  取值范围：[5，43200]，默认值为30。  单位：秒。  如果频道内无连麦方的状态持续超过该时间，录制程序会自动退出。退出后，再次调用start请求，会产生新的录制任务。  连麦方指：joiner或者publisher的用户。
	MaxIdleTime *int32 `json:"max_idle_time,omitempty"`

	RecordParam *RecordParam `json:"record_param,omitempty"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty"`

	// 任务状态。  - INIT：任务正在初始化 - RUNNING：任务正在运行 - STOPPED：任务已停止
	State *UpdateIndividualStreamJobResponseState `json:"state,omitempty"`

	// 任务结束原因 - TENANT_STOP - EXCEED_MAX_IDLE_TIME - INTERNAL_ERROR
	StopReason *UpdateIndividualStreamJobResponseStopReason `json:"stop_reason,omitempty"`

	// 针对任务状态的详细信息描述
	Description *string `json:"description,omitempty"`

	// 任务开始时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StartTime *string `json:"start_time,omitempty"`

	// 任务完成时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StopTime *string `json:"stop_time,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIndividualStreamJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndividualStreamJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateIndividualStreamJobResponse", string(data)}, " ")
}

type UpdateIndividualStreamJobResponseVideoType struct {
	value string
}

type UpdateIndividualStreamJobResponseVideoTypeEnum struct {
	CAMERASTREAM UpdateIndividualStreamJobResponseVideoType
	SCREENSTREAM UpdateIndividualStreamJobResponseVideoType
	EMPTY        UpdateIndividualStreamJobResponseVideoType
}

func GetUpdateIndividualStreamJobResponseVideoTypeEnum() UpdateIndividualStreamJobResponseVideoTypeEnum {
	return UpdateIndividualStreamJobResponseVideoTypeEnum{
		CAMERASTREAM: UpdateIndividualStreamJobResponseVideoType{
			value: "CAMERASTREAM",
		},
		SCREENSTREAM: UpdateIndividualStreamJobResponseVideoType{
			value: "SCREENSTREAM",
		},
		EMPTY: UpdateIndividualStreamJobResponseVideoType{
			value: "",
		},
	}
}

func (c UpdateIndividualStreamJobResponseVideoType) Value() string {
	return c.value
}

func (c UpdateIndividualStreamJobResponseVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIndividualStreamJobResponseVideoType) UnmarshalJSON(b []byte) error {
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

type UpdateIndividualStreamJobResponseSelectStreamType struct {
	value string
}

type UpdateIndividualStreamJobResponseSelectStreamTypeEnum struct {
	LD  UpdateIndividualStreamJobResponseSelectStreamType
	SD  UpdateIndividualStreamJobResponseSelectStreamType
	HD  UpdateIndividualStreamJobResponseSelectStreamType
	FHD UpdateIndividualStreamJobResponseSelectStreamType
}

func GetUpdateIndividualStreamJobResponseSelectStreamTypeEnum() UpdateIndividualStreamJobResponseSelectStreamTypeEnum {
	return UpdateIndividualStreamJobResponseSelectStreamTypeEnum{
		LD: UpdateIndividualStreamJobResponseSelectStreamType{
			value: "LD",
		},
		SD: UpdateIndividualStreamJobResponseSelectStreamType{
			value: "SD",
		},
		HD: UpdateIndividualStreamJobResponseSelectStreamType{
			value: "HD",
		},
		FHD: UpdateIndividualStreamJobResponseSelectStreamType{
			value: "FHD",
		},
	}
}

func (c UpdateIndividualStreamJobResponseSelectStreamType) Value() string {
	return c.value
}

func (c UpdateIndividualStreamJobResponseSelectStreamType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIndividualStreamJobResponseSelectStreamType) UnmarshalJSON(b []byte) error {
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

type UpdateIndividualStreamJobResponseState struct {
	value string
}

type UpdateIndividualStreamJobResponseStateEnum struct {
	INIT    UpdateIndividualStreamJobResponseState
	RUNNING UpdateIndividualStreamJobResponseState
	STOPPED UpdateIndividualStreamJobResponseState
}

func GetUpdateIndividualStreamJobResponseStateEnum() UpdateIndividualStreamJobResponseStateEnum {
	return UpdateIndividualStreamJobResponseStateEnum{
		INIT: UpdateIndividualStreamJobResponseState{
			value: "INIT",
		},
		RUNNING: UpdateIndividualStreamJobResponseState{
			value: "RUNNING",
		},
		STOPPED: UpdateIndividualStreamJobResponseState{
			value: "STOPPED",
		},
	}
}

func (c UpdateIndividualStreamJobResponseState) Value() string {
	return c.value
}

func (c UpdateIndividualStreamJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIndividualStreamJobResponseState) UnmarshalJSON(b []byte) error {
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

type UpdateIndividualStreamJobResponseStopReason struct {
	value string
}

type UpdateIndividualStreamJobResponseStopReasonEnum struct {
	TENANT_STOP          UpdateIndividualStreamJobResponseStopReason
	EXCEED_MAX_IDLE_TIME UpdateIndividualStreamJobResponseStopReason
	INTERNAL_ERROR       UpdateIndividualStreamJobResponseStopReason
}

func GetUpdateIndividualStreamJobResponseStopReasonEnum() UpdateIndividualStreamJobResponseStopReasonEnum {
	return UpdateIndividualStreamJobResponseStopReasonEnum{
		TENANT_STOP: UpdateIndividualStreamJobResponseStopReason{
			value: "TENANT_STOP",
		},
		EXCEED_MAX_IDLE_TIME: UpdateIndividualStreamJobResponseStopReason{
			value: "EXCEED_MAX_IDLE_TIME",
		},
		INTERNAL_ERROR: UpdateIndividualStreamJobResponseStopReason{
			value: "INTERNAL_ERROR",
		},
	}
}

func (c UpdateIndividualStreamJobResponseStopReason) Value() string {
	return c.value
}

func (c UpdateIndividualStreamJobResponseStopReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIndividualStreamJobResponseStopReason) UnmarshalJSON(b []byte) error {
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
