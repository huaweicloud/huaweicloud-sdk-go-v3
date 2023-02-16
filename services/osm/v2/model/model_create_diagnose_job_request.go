package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDiagnoseJobRequest struct {
	Body *SubmitDiagnoseJobReq `json:"body,omitempty"`
}

func (o CreateDiagnoseJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnoseJobRequest struct{}"
	}

	return strings.Join([]string{"CreateDiagnoseJobRequest", string(data)}, " ")
}
