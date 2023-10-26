package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCreateAudioStreamModerationJobRequest Request Object
type RunCreateAudioStreamModerationJobRequest struct {
	Body *AudioStreamCreateRequest `json:"body,omitempty"`
}

func (o RunCreateAudioStreamModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateAudioStreamModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunCreateAudioStreamModerationJobRequest", string(data)}, " ")
}
