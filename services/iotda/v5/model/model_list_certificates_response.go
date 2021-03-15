package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCertificatesResponse struct {
	// 证书列表。
	Certificates   *[]CertificatesRspDto `json:"certificates,omitempty"`
	Page           *Page                 `json:"page,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListCertificatesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListCertificatesResponse", string(data)}, " ")
}
