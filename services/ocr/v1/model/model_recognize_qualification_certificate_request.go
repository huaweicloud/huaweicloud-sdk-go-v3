package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizeQualificationCertificateRequest struct {
	Body *QualificationCertificateRequestBody `json:"body,omitempty"`
}

func (o RecognizeQualificationCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeQualificationCertificateRequest struct{}"
	}

	return strings.Join([]string{"RecognizeQualificationCertificateRequest", string(data)}, " ")
}
