package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
