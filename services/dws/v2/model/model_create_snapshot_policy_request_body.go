package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotPolicyRequestBody 创建备份策略信息
type CreateSnapshotPolicyRequestBody struct {

	// 保留天数。输入值必须在1-31之间。
	KeepDay *int32 `json:"keep_day,omitempty"`

	// 策略列表信息。当需要添加策略时该参数为必选。
	BackupStrategies *[]BackupStrategyRequest `json:"backup_strategies,omitempty"`

	// 备份设备。支持OBS、NBU和NFS。
	DeviceName *string `json:"device_name,omitempty"`

	// NBU备份介质端口。备份介质为NBU时该字段必填。
	ServerPort *string `json:"server_port,omitempty"`

	// NBU备份参数。备份介质为NBU时该字段必选。
	BackupParam *string `json:"backup_param,omitempty"`

	// 备份介质服务IP。备份介质为NBU和NFS时该字段必填。备份介质为NBU时表示NBU服务器地址，备份介质为NFS时表示NFS服务器地址。
	ServerIps *[]string `json:"server_ips,omitempty"`
}

func (o CreateSnapshotPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSnapshotPolicyRequestBody", string(data)}, " ")
}
