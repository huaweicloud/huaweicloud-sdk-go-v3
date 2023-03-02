package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateImageHighresolutionMattingTaskRequest struct {
	Body *ImageHighresolutionMattingRequestBody `json:"body,omitempty"`
}

func (o CreateImageHighresolutionMattingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageHighresolutionMattingTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateImageHighresolutionMattingTaskRequest", string(data)}, " ")
}
