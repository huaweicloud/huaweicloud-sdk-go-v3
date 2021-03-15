package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCertificatesResponse struct {
	// 证书列表，详情请参见CertificateDetail字段数据结构说明。
	Certificates *[]CertificateDetail `json:"certificates,omitempty"`
	// 证书数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCertificatesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListCertificatesResponse", string(data)}, " ")
}
