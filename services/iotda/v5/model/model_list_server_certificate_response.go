package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerCertificateResponse Response Object
type ListServerCertificateResponse struct {

	// **参数说明**：服务端证书列表。
	ServerCertificates *[]ServerCertificateDto `json:"server_certificates,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListServerCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerCertificateResponse struct{}"
	}

	return strings.Join([]string{"ListServerCertificateResponse", string(data)}, " ")
}
