package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 域名绑定和解绑证书的请求体
type AttachOrDetachCertsReqBody struct {

	// 证书的id集合
	CertificateIds []string `json:"certificate_ids"`
}

func (o AttachOrDetachCertsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachOrDetachCertsReqBody struct{}"
	}

	return strings.Join([]string{"AttachOrDetachCertsReqBody", string(data)}, " ")
}
