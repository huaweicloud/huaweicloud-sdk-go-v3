package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerHdaUpgradeRecord 服务器的accessAgent详细信息。
type ServerHdaUpgradeRecord struct {

	// 服务器id。
	ServerId *string `json:"server_id,omitempty"`

	// 机器名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 服务器名称。
	ServerName *string `json:"server_name,omitempty"`

	// 服务器组名称。
	ServerGroupName *string `json:"server_group_name,omitempty"`

	// 服务器的sid。
	Sid *string `json:"sid,omitempty"`

	// 当前的accessAgent版本。
	CurrentVersion *string `json:"current_version,omitempty"`

	// 目标的accessAgent版本。
	TargetVersion *string `json:"target_version,omitempty"`

	// HDA升级状态： - fail：升级失败 - upgrade success：升级成功
	UpgradeStatus *string `json:"upgrade_status,omitempty"`

	// 更新时间。
	UpgradeTime *string `json:"upgrade_time,omitempty"`
}

func (o ServerHdaUpgradeRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerHdaUpgradeRecord struct{}"
	}

	return strings.Join([]string{"ServerHdaUpgradeRecord", string(data)}, " ")
}
