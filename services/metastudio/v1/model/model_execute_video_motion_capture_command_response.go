package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteVideoMotionCaptureCommandResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteVideoMotionCaptureCommandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteVideoMotionCaptureCommandResponse struct{}"
	}

	return strings.Join([]string{"ExecuteVideoMotionCaptureCommandResponse", string(data)}, " ")
}
