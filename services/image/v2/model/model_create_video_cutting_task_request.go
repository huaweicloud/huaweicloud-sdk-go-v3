package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVideoCuttingTaskRequest struct {
	Body *VideoCuttingRequestBody `json:"body,omitempty"`
}

func (o CreateVideoCuttingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoCuttingTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoCuttingTaskRequest", string(data)}, " ")
}
