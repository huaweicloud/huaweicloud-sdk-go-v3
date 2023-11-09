package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnGatewayCertificateResponse Response Object
type ShowVpnGatewayCertificateResponse struct {
	Certificate *VpnGatewayCertificate `json:"certificate,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVpnGatewayCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnGatewayCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnGatewayCertificateResponse", string(data)}, " ")
}
