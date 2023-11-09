package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnGatewayCertificateRequest Request Object
type ShowVpnGatewayCertificateRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`
}

func (o ShowVpnGatewayCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnGatewayCertificateRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnGatewayCertificateRequest", string(data)}, " ")
}
