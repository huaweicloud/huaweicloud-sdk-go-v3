package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunMultiModalAssessmentRequest Request Object
type RunMultiModalAssessmentRequest struct {
	Body *PostMultiModalAssessmentReq `json:"body,omitempty"`
}

func (o RunMultiModalAssessmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunMultiModalAssessmentRequest struct{}"
	}

	return strings.Join([]string{"RunMultiModalAssessmentRequest", string(data)}, " ")
}
