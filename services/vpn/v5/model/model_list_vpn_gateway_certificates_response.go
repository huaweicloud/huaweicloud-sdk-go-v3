package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnGatewayCertificatesResponse Response Object
type ListVpnGatewayCertificatesResponse struct {

	// VPN网关证书信息
	VpnGatewayCertificates *[]VpnGatewayCertificate `json:"vpn_gateway_certificates,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVpnGatewayCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnGatewayCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListVpnGatewayCertificatesResponse", string(data)}, " ")
}
