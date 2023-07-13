package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSynthesisTaskRequest Request Object
type CreateSynthesisTaskRequest struct {
	Body *SynthesisTaskData `json:"body,omitempty"`
}

func (o CreateSynthesisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSynthesisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSynthesisTaskRequest", string(data)}, " ")
}
