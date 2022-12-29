package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FirmwareUpdateRecord struct {

	// 固件名称
	FirmwareName *string `json:"firmware_name,omitempty"`

	// 固件版本
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	// 固件大小
	FirmwareSize *int32 `json:"firmware_size,omitempty"`

	// 固件升级时间
	FirmwareUpgradeTime *string `json:"firmware_upgrade_time,omitempty"`

	// 固件升级状态
	Status *int32 `json:"status,omitempty"`
}

func (o FirmwareUpdateRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirmwareUpdateRecord struct{}"
	}

	return strings.Join([]string{"FirmwareUpdateRecord", string(data)}, " ")
}
