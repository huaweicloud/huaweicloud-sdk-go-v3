package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GatewayId 网关的ID。
type GatewayId struct {

	// 实例ID。
	GatewayId string `json:"gateway_id"`
}

func (o GatewayId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GatewayId struct{}"
	}

	return strings.Join([]string{"GatewayId", string(data)}, " ")
}
