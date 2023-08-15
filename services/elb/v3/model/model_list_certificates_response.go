package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesResponse Response Object
type ListCertificatesResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 证书对象列表。
	Certificates   *[]CertificateInfo `json:"certificates,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListCertificatesResponse", string(data)}, " ")
}
