package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotPolicyResp 自动快照信息
type SnapshotPolicyResp struct {

	// 集群是否开启自动快照。
	BackupEnable *bool `json:"backupEnable,omitempty"`

	// 快照备份时间。
	BakPeriod *string `json:"bakPeriod,omitempty"`

	// 快照备份间隔。
	BakFrequency *string `json:"bakFrequency,omitempty"`

	// 快照备份保留个数。
	BakKeepDay *int32 `json:"bakKeepDay,omitempty"`
}

func (o SnapshotPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotPolicyResp struct{}"
	}

	return strings.Join([]string{"SnapshotPolicyResp", string(data)}, " ")
}
