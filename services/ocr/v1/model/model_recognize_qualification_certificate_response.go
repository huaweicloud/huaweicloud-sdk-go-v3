package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeQualificationCertificateResponse struct {
	Result         *QualificationCertificateResult `json:"result,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o RecognizeQualificationCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeQualificationCertificateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeQualificationCertificateResponse", string(data)}, " ")
}
