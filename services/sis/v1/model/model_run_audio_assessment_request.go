package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunAudioAssessmentRequest Request Object
type RunAudioAssessmentRequest struct {
	Body *PostShortAudioAssessmentReq `json:"body,omitempty"`
}

func (o RunAudioAssessmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAudioAssessmentRequest struct{}"
	}

	return strings.Join([]string{"RunAudioAssessmentRequest", string(data)}, " ")
}
