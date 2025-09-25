package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotSettingReq struct {

	// 备份使用的OBS桶的桶名。
	Bucket string `json:"bucket"`

	// 委托名称，委托给CSS，允许CSS调用您的其他云服务。
	Agency string `json:"agency"`

	// 快照在OBS桶中的存放路径。
	BasePath *string `json:"base_path,omitempty"`

	// 配置每个节点的最大备份速率（每秒），即当备份的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际备份速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxSnapshotBytesPerSeconds *string `json:"max_snapshot_bytes_per_seconds,omitempty"`

	// 配置每个节点的最大恢复速率（每秒），即当恢复的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际恢复速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxRestoreBytesPerSeconds *string `json:"max_restore_bytes_per_seconds,omitempty"`

	// 是否开启自动创建快照策略。
	Enable *string `json:"enable,omitempty"`

	// 需要备份的索引名。
	Indices *string `json:"indices,omitempty"`

	// 自动创建快照的名称前缀，需要用户自己手动输入。
	Prefix *string `json:"prefix,omitempty"`

	// 每天创建快照的时刻。
	Period *string `json:"period,omitempty"`

	// 自定义设置快照保留的个数。系统在半点时刻会自动删除超过保留个数的快照。过期删除策略只针对与当前自动创建快照策略相同执行频次的自动快照。
	Keepday *int32 `json:"keepday,omitempty"`

	// 自动创建快照的执行频次。
	Frequency *string `json:"frequency,omitempty"`

	// 表示关闭自动创建快照策略时，是否需要清除所有自动创建的快照。
	DeleteAuto *string `json:"delete_auto,omitempty"`
}

func (o UpdateSnapshotSettingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotSettingReq struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotSettingReq", string(data)}, " ")
}
