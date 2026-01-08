package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertDetailRequest Request Object
type ShowCertDetailRequest struct {

	// 证书id。
	CertId string `json:"cert_id"`
}

func (o ShowCertDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCertDetailRequest", string(data)}, " ")
}
