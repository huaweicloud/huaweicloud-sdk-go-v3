package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FirmwareDetails 固件详细信息
type FirmwareDetails struct {

	// 组件
	ComponentName *string `json:"component_name,omitempty"`

	// 型号
	Model *string `json:"model,omitempty"`

	// 制造商
	Manufacturer *string `json:"manufacturer,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o FirmwareDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirmwareDetails struct{}"
	}

	return strings.Join([]string{"FirmwareDetails", string(data)}, " ")
}
