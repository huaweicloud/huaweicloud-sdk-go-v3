package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCertificateByCsrResponse struct {

	// 当前签发的证书ID。
	CertificateId  *string `json:"certificate_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCertificateByCsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateByCsrResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateByCsrResponse", string(data)}, " ")
}
