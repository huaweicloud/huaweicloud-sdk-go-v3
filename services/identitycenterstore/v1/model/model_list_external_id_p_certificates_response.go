package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExternalIdPCertificatesResponse Response Object
type ListExternalIdPCertificatesResponse struct {

	// 外部身份提供商证书列表
	IdpCertificates *[]IdpCertificate `json:"idp_certificates,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ListExternalIdPCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalIdPCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListExternalIdPCertificatesResponse", string(data)}, " ")
}
