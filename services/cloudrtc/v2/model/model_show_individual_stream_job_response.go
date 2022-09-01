package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowIndividualStreamJobResponse struct {

	// 任务编号
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 流名
	StreamName *string `json:"stream_name,omitempty" xml:"stream_name"`

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 房间id
	RoomId *string `json:"room_id,omitempty" xml:"room_id"`

	// 选看的用户id，单个录制任务内保证唯一
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	//  是否录制音频。  - true：录制音频 - false：不录制音频  缺省为true。
	IsRecordAudio *bool `json:"is_record_audio,omitempty" xml:"is_record_audio"`

	// 标识视频流的类型，可选摄像头流或者屏幕分享流，未填写表示不录制视频。  - CAMERASTREAM：摄像头视频流 - SCREENSTREAM：屏幕分享视频流  默认为CAMERASTREAM。
	VideoType *ShowIndividualStreamJobResponseVideoType `json:"video_type,omitempty" xml:"video_type"`

	// 指定窗口拉取的分辨率档位。  - LD - SD - HD - FHD  缺省为FHD。
	SelectStreamType *ShowIndividualStreamJobResponseSelectStreamType `json:"select_stream_type,omitempty" xml:"select_stream_type"`

	// 最长空闲频道时间。  取值范围：[5，43200]，默认值为30。  单位：秒。  如果频道内无连麦方的状态持续超过该时间，录制程序会自动退出。退出后，再次调用start请求，会产生新的录制任务。  连麦方指：joiner或者publisher的用户。
	MaxIdleTime *int32 `json:"max_idle_time,omitempty" xml:"max_idle_time"`

	RecordParam *RecordParam `json:"record_param,omitempty" xml:"record_param"`

	// 创建时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 任务状态。  - INIT：任务正在初始化 - RUNNING：任务正在运行 - STOPPED：任务已停止
	State *ShowIndividualStreamJobResponseState `json:"state,omitempty" xml:"state"`

	// 任务结束原因
	StopReason *ShowIndividualStreamJobResponseStopReason `json:"stop_reason,omitempty" xml:"stop_reason"`

	// 针对任务状态的详细信息描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 任务开始时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 任务完成时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIndividualStreamJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndividualStreamJobResponse struct{}"
	}

	return strings.Join([]string{"ShowIndividualStreamJobResponse", string(data)}, " ")
}

type ShowIndividualStreamJobResponseVideoType struct {
	value string
}

type ShowIndividualStreamJobResponseVideoTypeEnum struct {
	CAMERASTREAM ShowIndividualStreamJobResponseVideoType
	SCREENSTREAM ShowIndividualStreamJobResponseVideoType
	EMPTY        ShowIndividualStreamJobResponseVideoType
}

func GetShowIndividualStreamJobResponseVideoTypeEnum() ShowIndividualStreamJobResponseVideoTypeEnum {
	return ShowIndividualStreamJobResponseVideoTypeEnum{
		CAMERASTREAM: ShowIndividualStreamJobResponseVideoType{
			value: "CAMERASTREAM",
		},
		SCREENSTREAM: ShowIndividualStreamJobResponseVideoType{
			value: "SCREENSTREAM",
		},
		EMPTY: ShowIndividualStreamJobResponseVideoType{
			value: "",
		},
	}
}

func (c ShowIndividualStreamJobResponseVideoType) Value() string {
	return c.value
}

func (c ShowIndividualStreamJobResponseVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowIndividualStreamJobResponseVideoType) UnmarshalJSON(b []byte) error {
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

type ShowIndividualStreamJobResponseSelectStreamType struct {
	value string
}

type ShowIndividualStreamJobResponseSelectStreamTypeEnum struct {
	LD  ShowIndividualStreamJobResponseSelectStreamType
	SD  ShowIndividualStreamJobResponseSelectStreamType
	HD  ShowIndividualStreamJobResponseSelectStreamType
	FHD ShowIndividualStreamJobResponseSelectStreamType
}

func GetShowIndividualStreamJobResponseSelectStreamTypeEnum() ShowIndividualStreamJobResponseSelectStreamTypeEnum {
	return ShowIndividualStreamJobResponseSelectStreamTypeEnum{
		LD: ShowIndividualStreamJobResponseSelectStreamType{
			value: "LD",
		},
		SD: ShowIndividualStreamJobResponseSelectStreamType{
			value: "SD",
		},
		HD: ShowIndividualStreamJobResponseSelectStreamType{
			value: "HD",
		},
		FHD: ShowIndividualStreamJobResponseSelectStreamType{
			value: "FHD",
		},
	}
}

func (c ShowIndividualStreamJobResponseSelectStreamType) Value() string {
	return c.value
}

func (c ShowIndividualStreamJobResponseSelectStreamType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowIndividualStreamJobResponseSelectStreamType) UnmarshalJSON(b []byte) error {
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

type ShowIndividualStreamJobResponseState struct {
	value string
}

type ShowIndividualStreamJobResponseStateEnum struct {
	INIT    ShowIndividualStreamJobResponseState
	RUNNING ShowIndividualStreamJobResponseState
	STOPPED ShowIndividualStreamJobResponseState
}

func GetShowIndividualStreamJobResponseStateEnum() ShowIndividualStreamJobResponseStateEnum {
	return ShowIndividualStreamJobResponseStateEnum{
		INIT: ShowIndividualStreamJobResponseState{
			value: "INIT",
		},
		RUNNING: ShowIndividualStreamJobResponseState{
			value: "RUNNING",
		},
		STOPPED: ShowIndividualStreamJobResponseState{
			value: "STOPPED",
		},
	}
}

func (c ShowIndividualStreamJobResponseState) Value() string {
	return c.value
}

func (c ShowIndividualStreamJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowIndividualStreamJobResponseState) UnmarshalJSON(b []byte) error {
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

type ShowIndividualStreamJobResponseStopReason struct {
	value string
}

type ShowIndividualStreamJobResponseStopReasonEnum struct {
	TENANT_STOP          ShowIndividualStreamJobResponseStopReason
	EXCEED_MAX_IDLE_TIME ShowIndividualStreamJobResponseStopReason
	INTERNAL_ERROR       ShowIndividualStreamJobResponseStopReason
}

func GetShowIndividualStreamJobResponseStopReasonEnum() ShowIndividualStreamJobResponseStopReasonEnum {
	return ShowIndividualStreamJobResponseStopReasonEnum{
		TENANT_STOP: ShowIndividualStreamJobResponseStopReason{
			value: "TENANT_STOP",
		},
		EXCEED_MAX_IDLE_TIME: ShowIndividualStreamJobResponseStopReason{
			value: "EXCEED_MAX_IDLE_TIME",
		},
		INTERNAL_ERROR: ShowIndividualStreamJobResponseStopReason{
			value: "INTERNAL_ERROR",
		},
	}
}

func (c ShowIndividualStreamJobResponseStopReason) Value() string {
	return c.value
}

func (c ShowIndividualStreamJobResponseStopReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowIndividualStreamJobResponseStopReason) UnmarshalJSON(b []byte) error {
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
