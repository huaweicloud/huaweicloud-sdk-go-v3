package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificatePrivateKeyEchoRequestBody This is a auto create Body Object
type CreateCertificatePrivateKeyEchoRequestBody struct {

	// 证书回显开关，项目粒度的,默认情况下,\"private_key_echo\"是true,证书的返回体中展示私钥。 当值为false时,证书的返回体中不展示私钥。
	PrivateKeyEcho bool `json:"private_key_echo"`
}

func (o CreateCertificatePrivateKeyEchoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificatePrivateKeyEchoRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificatePrivateKeyEchoRequestBody", string(data)}, " ")
}
