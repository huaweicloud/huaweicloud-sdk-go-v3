package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferPolicy 转储策略信息
type TransferPolicy struct {

	// 策略id
	SettingId *string `json:"setting_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 转储策略周期
	Period *string `json:"period,omitempty"`

	// 转储备份类型
	BackupType *string `json:"backup_type,omitempty"`

	// 目标存储
	Destination *string `json:"destination,omitempty"`

	// 转储目标前缀
	Prefix *string `json:"prefix,omitempty"`
}

func (o TransferPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferPolicy struct{}"
	}

	return strings.Join([]string{"TransferPolicy", string(data)}, " ")
}
