package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}
