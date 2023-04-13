package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 视频驱动动作任务请求。
type VideoMotionCaptureJobReq struct {

	// 视频驱动模式。 * HEAD: 头部 * HALF_BODY: 半身 * FULL_BODY: 全身 * AUTO: 自动
	MotionCaptureMode *VideoMotionCaptureJobReqMotionCaptureMode `json:"motion_capture_mode,omitempty"`

	InputInfo *InputInfo `json:"input_info,omitempty"`

	OutputInfo *OutputInfo `json:"output_info,omitempty"`
}

func (o VideoMotionCaptureJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoMotionCaptureJobReq struct{}"
	}

	return strings.Join([]string{"VideoMotionCaptureJobReq", string(data)}, " ")
}

type VideoMotionCaptureJobReqMotionCaptureMode struct {
	value string
}

type VideoMotionCaptureJobReqMotionCaptureModeEnum struct {
	HEAD      VideoMotionCaptureJobReqMotionCaptureMode
	HALF_BODY VideoMotionCaptureJobReqMotionCaptureMode
	FULL_BODY VideoMotionCaptureJobReqMotionCaptureMode
	AUTO      VideoMotionCaptureJobReqMotionCaptureMode
}

func GetVideoMotionCaptureJobReqMotionCaptureModeEnum() VideoMotionCaptureJobReqMotionCaptureModeEnum {
	return VideoMotionCaptureJobReqMotionCaptureModeEnum{
		HEAD: VideoMotionCaptureJobReqMotionCaptureMode{
			value: "HEAD",
		},
		HALF_BODY: VideoMotionCaptureJobReqMotionCaptureMode{
			value: "HALF_BODY",
		},
		FULL_BODY: VideoMotionCaptureJobReqMotionCaptureMode{
			value: "FULL_BODY",
		},
		AUTO: VideoMotionCaptureJobReqMotionCaptureMode{
			value: "AUTO",
		},
	}
}

func (c VideoMotionCaptureJobReqMotionCaptureMode) Value() string {
	return c.value
}

func (c VideoMotionCaptureJobReqMotionCaptureMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoMotionCaptureJobReqMotionCaptureMode) UnmarshalJSON(b []byte) error {
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
