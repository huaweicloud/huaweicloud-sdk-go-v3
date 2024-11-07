package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GatewayType 网关的类型。
type GatewayType struct {
	GatewayType *GatewayTypeEnum `json:"gateway_type"`
}

func (o GatewayType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GatewayType struct{}"
	}

	return strings.Join([]string{"GatewayType", string(data)}, " ")
}
