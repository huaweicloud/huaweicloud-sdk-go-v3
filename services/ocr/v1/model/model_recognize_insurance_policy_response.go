package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeInsurancePolicyResponse Response Object
type RecognizeInsurancePolicyResponse struct {
	Result *InsurancePolicyResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeInsurancePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeInsurancePolicyResponse struct{}"
	}

	return strings.Join([]string{"RecognizeInsurancePolicyResponse", string(data)}, " ")
}
