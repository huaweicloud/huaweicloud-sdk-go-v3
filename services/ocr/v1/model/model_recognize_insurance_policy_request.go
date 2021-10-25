package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeInsurancePolicyRequest struct {
	Body *InsurancePolicyRequestBody `json:"body,omitempty"`
}

func (o RecognizeInsurancePolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeInsurancePolicyRequest struct{}"
	}

	return strings.Join([]string{"RecognizeInsurancePolicyRequest", string(data)}, " ")
}
