package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCertificateResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Certificate    *CertificateInfo `json:"certificate,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateResponse", string(data)}, " ")
}
