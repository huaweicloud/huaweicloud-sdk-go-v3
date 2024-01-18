package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnGatewayCertificateRequestBody struct {
	Certificate *VpnGatewayCertificateRequestBodyContent `json:"certificate"`
}

func (o VpnGatewayCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnGatewayCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"VpnGatewayCertificateRequestBody", string(data)}, " ")
}
