package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunCreateAudioModerationJobRequest struct {
	Body *AudioCreateRequest `json:"body,omitempty" xml:"body"`
}

func (o RunCreateAudioModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateAudioModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunCreateAudioModerationJobRequest", string(data)}, " ")
}
