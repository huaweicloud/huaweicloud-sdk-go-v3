package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVideoCoverAnalysisTaskRequest struct {
	Body *VideoCoverAnalysisCreateTaskRequestBody `json:"body,omitempty"`
}

func (o CreateVideoCoverAnalysisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoCoverAnalysisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoCoverAnalysisTaskRequest", string(data)}, " ")
}
