package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertificateAuthorityRequest Request Object
type DeleteCertificateAuthorityRequest struct {

	// 所要计划删除的CA证书ID。
	CaId string `json:"ca_id"`

	// 延迟删除时间，单位为”天“。
	PendingDays string `json:"pending_days"`
}

func (o DeleteCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"DeleteCertificateAuthorityRequest", string(data)}, " ")
}
