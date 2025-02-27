package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeGateway struct {

	// 网关ID，用于标识设备所属的父设备，即父设备的设备ID。
	GatewayId string `json:"gateway_id"`
}

func (o ChangeGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeGateway struct{}"
	}

	return strings.Join([]string{"ChangeGateway", string(data)}, " ")
}
