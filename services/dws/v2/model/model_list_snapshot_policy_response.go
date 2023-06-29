package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotPolicyResponse Response Object
type ListSnapshotPolicyResponse struct {

	// 保留天数。
	KeepDay *string `json:"keep_day,omitempty"`

	// 备份策略。
	BackupStrategies *[]BackupStrategyDetail `json:"backup_strategies,omitempty"`

	// 备份设备。
	DeviceName *string `json:"device_name,omitempty"`

	// 服务IP。
	ServerIps *[]string `json:"server_ips,omitempty"`

	// 服务端口。
	ServerPort *string `json:"server_port,omitempty"`

	// 参数。
	BackupParam    *string `json:"backup_param,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotPolicyResponse", string(data)}, " ")
}
