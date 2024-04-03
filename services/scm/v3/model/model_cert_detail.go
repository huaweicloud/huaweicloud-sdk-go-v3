package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertDetail struct {

	// 证书ID。
	CertId string `json:"cert_id"`
}

func (o CertDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertDetail struct{}"
	}

	return strings.Join([]string{"CertDetail", string(data)}, " ")
}
