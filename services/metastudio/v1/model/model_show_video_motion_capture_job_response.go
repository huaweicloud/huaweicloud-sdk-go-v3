package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowVideoMotionCaptureJobResponse Response Object
type ShowVideoMotionCaptureJobResponse struct {

	// 视频驱动模式。 * HEAD：头部 * HALF_BODY：半身 * FULL_BODY：全身 * AUTO：自动
	MotionCaptureMode *ShowVideoMotionCaptureJobResponseMotionCaptureMode `json:"motion_capture_mode,omitempty"`

	InputInfo *InputInfo `json:"input_info,omitempty"`

	OutputInfo *OutputInfo `json:"output_info,omitempty"`

	// 视频驱动任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务的状态。 * WAITING：等待中 * PROCESSING：处理中 * SUCCEED：成功 * FAILED：失败
	State *ShowVideoMotionCaptureJobResponseState `json:"state,omitempty"`

	// 任务开始时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，格式遵循：RFC 3339。 例 “2020-07-30T10:43:17Z”。
	EndTime *string `json:"end_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVideoMotionCaptureJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoMotionCaptureJobResponse struct{}"
	}

	return strings.Join([]string{"ShowVideoMotionCaptureJobResponse", string(data)}, " ")
}

type ShowVideoMotionCaptureJobResponseMotionCaptureMode struct {
	value string
}

type ShowVideoMotionCaptureJobResponseMotionCaptureModeEnum struct {
	HEAD      ShowVideoMotionCaptureJobResponseMotionCaptureMode
	HALF_BODY ShowVideoMotionCaptureJobResponseMotionCaptureMode
	FULL_BODY ShowVideoMotionCaptureJobResponseMotionCaptureMode
	AUTO      ShowVideoMotionCaptureJobResponseMotionCaptureMode
}

func GetShowVideoMotionCaptureJobResponseMotionCaptureModeEnum() ShowVideoMotionCaptureJobResponseMotionCaptureModeEnum {
	return ShowVideoMotionCaptureJobResponseMotionCaptureModeEnum{
		HEAD: ShowVideoMotionCaptureJobResponseMotionCaptureMode{
			value: "HEAD",
		},
		HALF_BODY: ShowVideoMotionCaptureJobResponseMotionCaptureMode{
			value: "HALF_BODY",
		},
		FULL_BODY: ShowVideoMotionCaptureJobResponseMotionCaptureMode{
			value: "FULL_BODY",
		},
		AUTO: ShowVideoMotionCaptureJobResponseMotionCaptureMode{
			value: "AUTO",
		},
	}
}

func (c ShowVideoMotionCaptureJobResponseMotionCaptureMode) Value() string {
	return c.value
}

func (c ShowVideoMotionCaptureJobResponseMotionCaptureMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVideoMotionCaptureJobResponseMotionCaptureMode) UnmarshalJSON(b []byte) error {
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

type ShowVideoMotionCaptureJobResponseState struct {
	value string
}

type ShowVideoMotionCaptureJobResponseStateEnum struct {
	WAITING    ShowVideoMotionCaptureJobResponseState
	PROCESSING ShowVideoMotionCaptureJobResponseState
	SUCCEED    ShowVideoMotionCaptureJobResponseState
	FAILED     ShowVideoMotionCaptureJobResponseState
}

func GetShowVideoMotionCaptureJobResponseStateEnum() ShowVideoMotionCaptureJobResponseStateEnum {
	return ShowVideoMotionCaptureJobResponseStateEnum{
		WAITING: ShowVideoMotionCaptureJobResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowVideoMotionCaptureJobResponseState{
			value: "PROCESSING",
		},
		SUCCEED: ShowVideoMotionCaptureJobResponseState{
			value: "SUCCEED",
		},
		FAILED: ShowVideoMotionCaptureJobResponseState{
			value: "FAILED",
		},
	}
}

func (c ShowVideoMotionCaptureJobResponseState) Value() string {
	return c.value
}

func (c ShowVideoMotionCaptureJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVideoMotionCaptureJobResponseState) UnmarshalJSON(b []byte) error {
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
