package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSslCertDownloadAddressResponse Response Object
type ListSslCertDownloadAddressResponse struct {

	// 证书列表
	Certs          *[]CertInfo `json:"certs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListSslCertDownloadAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSslCertDownloadAddressResponse struct{}"
	}

	return strings.Join([]string{"ListSslCertDownloadAddressResponse", string(data)}, " ")
}
