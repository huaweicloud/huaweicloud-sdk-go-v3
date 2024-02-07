package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnGatewayCertificateRequestBody struct {
	Certificate *CreateVpnGatewayCertificateRequestBodyContent `json:"certificate"`
}

func (o CreateVpnGatewayCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnGatewayCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpnGatewayCertificateRequestBody", string(data)}, " ")
}
