package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClientIpTransparentTransmissionRequestBody struct {

	// 开启或关闭客户ip透传标志
	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`
}

func (o UpdateClientIpTransparentTransmissionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientIpTransparentTransmissionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateClientIpTransparentTransmissionRequestBody", string(data)}, " ")
}
