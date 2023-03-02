package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVideoSynthesisTaskRequest struct {
	Body *VideoSynthesisRequestBody `json:"body,omitempty"`
}

func (o CreateVideoSynthesisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoSynthesisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoSynthesisTaskRequest", string(data)}, " ")
}
