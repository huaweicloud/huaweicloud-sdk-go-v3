package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeQualificationCertificateResponse Response Object
type RecognizeQualificationCertificateResponse struct {
	Result *QualificationCertificateResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeQualificationCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeQualificationCertificateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeQualificationCertificateResponse", string(data)}, " ")
}
