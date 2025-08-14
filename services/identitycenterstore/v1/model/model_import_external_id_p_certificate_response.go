package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportExternalIdPCertificateResponse Response Object
type ImportExternalIdPCertificateResponse struct {

	// 证书的全局唯一标识符（ID）
	CertificateId  *string `json:"certificate_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportExternalIdPCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportExternalIdPCertificateResponse struct{}"
	}

	return strings.Join([]string{"ImportExternalIdPCertificateResponse", string(data)}, " ")
}
