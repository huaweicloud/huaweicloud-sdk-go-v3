package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeInsurancePolicyResponse struct {
	Result         *InsurancePolicyResult `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o RecognizeInsurancePolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeInsurancePolicyResponse struct{}"
	}

	return strings.Join([]string{"RecognizeInsurancePolicyResponse", string(data)}, " ")
}
