package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCreateVideoStreamModerationJobRequest Request Object
type RunCreateVideoStreamModerationJobRequest struct {
	Body *VideoStreamCreateRequest `json:"body,omitempty"`
}

func (o RunCreateVideoStreamModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateVideoStreamModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunCreateVideoStreamModerationJobRequest", string(data)}, " ")
}
