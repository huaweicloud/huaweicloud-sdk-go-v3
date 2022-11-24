package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportCertificateAuthorityCsrRequest struct {

	// 未激活的从属CA证书ID。
	CaId string `json:"ca_id"`
}

func (o ExportCertificateAuthorityCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateAuthorityCsrRequest struct{}"
	}

	return strings.Join([]string{"ExportCertificateAuthorityCsrRequest", string(data)}, " ")
}
