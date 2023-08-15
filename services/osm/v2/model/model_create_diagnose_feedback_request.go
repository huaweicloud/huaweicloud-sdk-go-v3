package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiagnoseFeedbackRequest Request Object
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
