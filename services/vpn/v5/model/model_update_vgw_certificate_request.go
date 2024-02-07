package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVgwCertificateRequest Request Object
type UpdateVgwCertificateRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`

	// VPN网关证书ID
	CertificateId string `json:"certificate_id"`

	Body *UpdateVpnGatewayCertificateRequestBody `json:"body,omitempty"`
}

func (o UpdateVgwCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwCertificateRequest struct{}"
	}

	return strings.Join([]string{"UpdateVgwCertificateRequest", string(data)}, " ")
}
