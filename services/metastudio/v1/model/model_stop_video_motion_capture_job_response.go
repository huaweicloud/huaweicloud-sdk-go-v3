package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopVideoMotionCaptureJobResponse Response Object
type StopVideoMotionCaptureJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopVideoMotionCaptureJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopVideoMotionCaptureJobResponse struct{}"
	}

	return strings.Join([]string{"StopVideoMotionCaptureJobResponse", string(data)}, " ")
}
