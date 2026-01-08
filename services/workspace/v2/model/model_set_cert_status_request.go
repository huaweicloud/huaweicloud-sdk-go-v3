package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetCertStatusRequest Request Object
type SetCertStatusRequest struct {

	// 证书id。
	CertId string `json:"cert_id"`

	// 证书状态, enable启用, disable禁用。
	Status string `json:"status"`
}

func (o SetCertStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCertStatusRequest struct{}"
	}

	return strings.Join([]string{"SetCertStatusRequest", string(data)}, " ")
}
