package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCreateAudioModerationJobResponse Response Object
type RunCreateAudioModerationJobResponse struct {

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId *string `json:"request_id,omitempty"`

	// 作业唯一标识。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunCreateAudioModerationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateAudioModerationJobResponse struct{}"
	}

	return strings.Join([]string{"RunCreateAudioModerationJobResponse", string(data)}, " ")
}
