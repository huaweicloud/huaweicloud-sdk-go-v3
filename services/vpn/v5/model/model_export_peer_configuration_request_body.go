package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportPeerConfigurationRequestBody struct {

	// 设备厂商
	Vendor string `json:"vendor"`

	// 设备系列
	Type string `json:"type"`

	// 设备型号
	Model string `json:"model"`

	// 设备版本
	Version string `json:"version"`
}

func (o ExportPeerConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportPeerConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ExportPeerConfigurationRequestBody", string(data)}, " ")
}
