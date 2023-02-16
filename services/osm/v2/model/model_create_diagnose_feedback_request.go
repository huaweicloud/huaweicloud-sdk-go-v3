package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDiagnoseFeedbackRequest struct {
	Body *CreateFeedbackReq `json:"body,omitempty"`
}

func (o CreateDiagnoseFeedbackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnoseFeedbackRequest struct{}"
	}

	return strings.Join([]string{"CreateDiagnoseFeedbackRequest", string(data)}, " ")
}
