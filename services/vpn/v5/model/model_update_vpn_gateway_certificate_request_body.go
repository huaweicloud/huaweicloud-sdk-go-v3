package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnGatewayCertificateRequestBody struct {
	Certificate *UpdateVpnGatewayCertificateRequestBodyContent `json:"certificate"`
}

func (o UpdateVpnGatewayCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnGatewayCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpnGatewayCertificateRequestBody", string(data)}, " ")
}
