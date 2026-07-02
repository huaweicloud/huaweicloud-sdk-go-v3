package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupportedDevice struct {

	// 厂商
	Vendor *string `json:"vendor,omitempty"`

	// 型号
	Platform *string `json:"platform,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o SupportedDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportedDevice struct{}"
	}

	return strings.Join([]string{"SupportedDevice", string(data)}, " ")
}
