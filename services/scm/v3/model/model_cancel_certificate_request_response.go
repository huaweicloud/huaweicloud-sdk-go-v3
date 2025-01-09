package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelCertificateRequestResponse Response Object
type CancelCertificateRequestResponse struct {

	// 证书ID。
	CertId *string `json:"cert_id,omitempty"`

	// 执行结果。取值如下： success：申请成功。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelCertificateRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelCertificateRequestResponse struct{}"
	}

	return strings.Join([]string{"CancelCertificateRequestResponse", string(data)}, " ")
}
