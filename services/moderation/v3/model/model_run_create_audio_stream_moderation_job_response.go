package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCreateAudioStreamModerationJobResponse Response Object
type RunCreateAudioStreamModerationJobResponse struct {

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId *string `json:"request_id,omitempty"`

	// 作业唯一标识。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunCreateAudioStreamModerationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateAudioStreamModerationJobResponse struct{}"
	}

	return strings.Join([]string{"RunCreateAudioStreamModerationJobResponse", string(data)}, " ")
}
