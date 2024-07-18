package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificatePrivateKeyEchoRequest Request Object
type CreateCertificatePrivateKeyEchoRequest struct {
	Body *CreateCertificatePrivateKeyEchoRequestBody `json:"body,omitempty"`
}

func (o CreateCertificatePrivateKeyEchoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificatePrivateKeyEchoRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificatePrivateKeyEchoRequest", string(data)}, " ")
}
