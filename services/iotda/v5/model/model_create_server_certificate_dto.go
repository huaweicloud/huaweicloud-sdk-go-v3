package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerCertificateDto 创建服务器证书请求体。
type CreateServerCertificateDto struct {

	// **参数说明**：证书内容，PEM格式
	CertificatePem string `json:"certificate_pem"`

	// **参数说明**：证书私钥
	PrivateKey string `json:"private_key"`

	// **参数说明**：私钥密码
	PrivateKeyPassword *string `json:"private_key_password,omitempty"`
}

func (o CreateServerCertificateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerCertificateDto struct{}"
	}

	return strings.Join([]string{"CreateServerCertificateDto", string(data)}, " ")
}
