package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParseCertificateSigningRequestRequestBody struct {

	// 证书签名请求。请使用“\\r\\n”或“\\n”替代证书签名请求中的换行符，若通过console端请求此接口，则无需做符号转换。
	Csr string `json:"csr"`
}

func (o ParseCertificateSigningRequestRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseCertificateSigningRequestRequestBody struct{}"
	}

	return strings.Join([]string{"ParseCertificateSigningRequestRequestBody", string(data)}, " ")
}
