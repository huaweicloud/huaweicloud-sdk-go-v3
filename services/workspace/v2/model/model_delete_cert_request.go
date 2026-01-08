package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertRequest Request Object
type DeleteCertRequest struct {

	// 证书id。
	CertId string `json:"cert_id"`
}

func (o DeleteCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertRequest struct{}"
	}

	return strings.Join([]string{"DeleteCertRequest", string(data)}, " ")
}
