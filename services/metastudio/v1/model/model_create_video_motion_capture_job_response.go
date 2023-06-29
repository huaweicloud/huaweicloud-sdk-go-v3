package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVideoMotionCaptureJobResponse Response Object
type CreateVideoMotionCaptureJobResponse struct {

	// 视频驱动动作任务ID
	JobId *string `json:"job_id,omitempty"`

	RtcRoomInfo    *RtcRoomInfoList `json:"rtc_room_info,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateVideoMotionCaptureJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoMotionCaptureJobResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoMotionCaptureJobResponse", string(data)}, " ")
}
