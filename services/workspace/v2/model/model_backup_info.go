package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupInfo 备份信息。
type BackupInfo struct {

	// 优先级，数字越小，优先级越高，取值1-255。
	Priority int32 `json:"priority"`

	// 接入备份地址。
	Address string `json:"address"`

	// 租户自定义接入备份地址。
	AddressCustom *string `json:"address_custom,omitempty"`
}

func (o BackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInfo struct{}"
	}

	return strings.Join([]string{"BackupInfo", string(data)}, " ")
}
