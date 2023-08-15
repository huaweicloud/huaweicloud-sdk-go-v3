package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupPolicy 备份策略
type BackupPolicy struct {

	// 保留天数。
	KeepDay *int32 `json:"keep_day,omitempty"`

	BackupStrategies *BackupStrategyDetail `json:"backup_strategies,omitempty"`

	// 备份设备。
	DeviceName *string `json:"device_name,omitempty"`

	// 端口。
	ServerPort *string `json:"server_port,omitempty"`

	// 参数。
	BackupParam *string `json:"backup_param,omitempty"`

	// 备份介质服务IP。
	ServerIps *[]string `json:"server_ips,omitempty"`
}

func (o BackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicy struct{}"
	}

	return strings.Join([]string{"BackupPolicy", string(data)}, " ")
}
