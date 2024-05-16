package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeRealEstateCertificateResponse Response Object
type RecognizeRealEstateCertificateResponse struct {
	Result *RealEstateCertificateResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeRealEstateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeRealEstateCertificateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeRealEstateCertificateResponse", string(data)}, " ")
}
