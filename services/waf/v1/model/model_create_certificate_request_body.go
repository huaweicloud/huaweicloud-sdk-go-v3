package model

import (
	"encoding/json"

	"strings"
)

type CreateCertificateRequestBody struct {
	// 证书名

	Name *string `json:"name,omitempty"`
	// 证书文件，PEM编码

	Content *string `json:"content,omitempty"`
	// 证书私钥，PEM编码

	Key *string `json:"key,omitempty"`
}

func (o CreateCertificateRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}
