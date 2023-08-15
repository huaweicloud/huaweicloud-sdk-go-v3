package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunModerationAudioRequest Request Object
type RunModerationAudioRequest struct {
	Body *RunModerationAudioRequestBody `json:"body,omitempty"`
}

func (o RunModerationAudioRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioRequest struct{}"
	}

	return strings.Join([]string{"RunModerationAudioRequest", string(data)}, " ")
}
