package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClientCertTrustedCerts struct {

	// **参数解释：** 客户端CA证书的内容 **约束限制：** 仅支持PEM格式 **取值范围：** 不涉及
	TrustedCert *string `json:"trusted_cert,omitempty"`
}

func (o ClientCertTrustedCerts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientCertTrustedCerts struct{}"
	}

	return strings.Join([]string{"ClientCertTrustedCerts", string(data)}, " ")
}
