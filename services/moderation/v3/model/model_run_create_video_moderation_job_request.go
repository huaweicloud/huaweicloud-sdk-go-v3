package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCreateVideoModerationJobRequest Request Object
type RunCreateVideoModerationJobRequest struct {
	Body *VideoCreateRequest `json:"body,omitempty"`
}

func (o RunCreateVideoModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateVideoModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunCreateVideoModerationJobRequest", string(data)}, " ")
}
