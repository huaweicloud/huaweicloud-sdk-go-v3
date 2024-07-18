package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificatePrivateKeyEchoRequest Request Object
type ShowCertificatePrivateKeyEchoRequest struct {
}

func (o ShowCertificatePrivateKeyEchoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificatePrivateKeyEchoRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificatePrivateKeyEchoRequest", string(data)}, " ")
}
