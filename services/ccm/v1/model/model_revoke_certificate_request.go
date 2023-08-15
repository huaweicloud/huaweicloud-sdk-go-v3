package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeCertificateRequest Request Object
type RevokeCertificateRequest struct {

	// 所要吊销的私有证书ID。
	CertificateId string `json:"certificate_id"`

	Body *RevokeCertificateRequestBody `json:"body,omitempty"`
}

func (o RevokeCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeCertificateRequest struct{}"
	}

	return strings.Join([]string{"RevokeCertificateRequest", string(data)}, " ")
}
