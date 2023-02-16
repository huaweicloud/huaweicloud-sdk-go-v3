package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateFeedbackRequest struct {
	Body *CreateFeedbackV2Req `json:"body,omitempty"`
}

func (o CreateFeedbackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFeedbackRequest struct{}"
	}

	return strings.Join([]string{"CreateFeedbackRequest", string(data)}, " ")
}
