package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeInsurancePolicyRequest struct {
	Body *InsurancePolicyRequestBody `json:"body,omitempty"`
}

func (o RecognizeInsurancePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeInsurancePolicyRequest struct{}"
	}

	return strings.Join([]string{"RecognizeInsurancePolicyRequest", string(data)}, " ")
}
