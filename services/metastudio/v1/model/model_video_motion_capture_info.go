package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 视频驱动任务详情。
type VideoMotionCaptureInfo struct {

	// 视频驱动模式。 * HEAD: 头部 * HALF_BODY: 半身 * FULL_BODY: 全身 * AUTO: 自动
	MotionCaptureMode *VideoMotionCaptureInfoMotionCaptureMode `json:"motion_capture_mode,omitempty"`

	InputInfo *InputInfo `json:"input_info,omitempty"`

	OutputInfo *OutputInfo `json:"output_info,omitempty"`

	// 视频驱动任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务的状态。 * WAITING: 等待中 * PROCESSING: 处理中 * SUCCEED: 成功 * FAILED: 失败
	State *VideoMotionCaptureInfoState `json:"state,omitempty"`

	// 任务开始时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间,格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	EndTime *string `json:"end_time,omitempty"`
}

func (o VideoMotionCaptureInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoMotionCaptureInfo struct{}"
	}

	return strings.Join([]string{"VideoMotionCaptureInfo", string(data)}, " ")
}

type VideoMotionCaptureInfoMotionCaptureMode struct {
	value string
}

type VideoMotionCaptureInfoMotionCaptureModeEnum struct {
	HEAD      VideoMotionCaptureInfoMotionCaptureMode
	HALF_BODY VideoMotionCaptureInfoMotionCaptureMode
	FULL_BODY VideoMotionCaptureInfoMotionCaptureMode
	AUTO      VideoMotionCaptureInfoMotionCaptureMode
}

func GetVideoMotionCaptureInfoMotionCaptureModeEnum() VideoMotionCaptureInfoMotionCaptureModeEnum {
	return VideoMotionCaptureInfoMotionCaptureModeEnum{
		HEAD: VideoMotionCaptureInfoMotionCaptureMode{
			value: "HEAD",
		},
		HALF_BODY: VideoMotionCaptureInfoMotionCaptureMode{
			value: "HALF_BODY",
		},
		FULL_BODY: VideoMotionCaptureInfoMotionCaptureMode{
			value: "FULL_BODY",
		},
		AUTO: VideoMotionCaptureInfoMotionCaptureMode{
			value: "AUTO",
		},
	}
}

func (c VideoMotionCaptureInfoMotionCaptureMode) Value() string {
	return c.value
}

func (c VideoMotionCaptureInfoMotionCaptureMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoMotionCaptureInfoMotionCaptureMode) UnmarshalJSON(b []byte) error {
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

type VideoMotionCaptureInfoState struct {
	value string
}

type VideoMotionCaptureInfoStateEnum struct {
	WAITING    VideoMotionCaptureInfoState
	PROCESSING VideoMotionCaptureInfoState
	SUCCEED    VideoMotionCaptureInfoState
	FAILED     VideoMotionCaptureInfoState
}

func GetVideoMotionCaptureInfoStateEnum() VideoMotionCaptureInfoStateEnum {
	return VideoMotionCaptureInfoStateEnum{
		WAITING: VideoMotionCaptureInfoState{
			value: "WAITING",
		},
		PROCESSING: VideoMotionCaptureInfoState{
			value: "PROCESSING",
		},
		SUCCEED: VideoMotionCaptureInfoState{
			value: "SUCCEED",
		},
		FAILED: VideoMotionCaptureInfoState{
			value: "FAILED",
		},
	}
}

func (c VideoMotionCaptureInfoState) Value() string {
	return c.value
}

func (c VideoMotionCaptureInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoMotionCaptureInfoState) UnmarshalJSON(b []byte) error {
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
