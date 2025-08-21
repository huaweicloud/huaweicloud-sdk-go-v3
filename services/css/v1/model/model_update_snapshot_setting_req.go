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
	BasePath string `json:"base_path"`

	// 配置每个节点的最大备份速率（每秒），即当备份的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际备份速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxSnapshotBytesPerSeconds *string `json:"maxSnapshotBytesPerSeconds,omitempty"`

	// 配置每个节点的最大恢复速率（每秒），即当恢复的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际恢复速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxRestoreBytesPerSeconds *string `json:"maxRestoreBytesPerSeconds,omitempty"`
}

func (o UpdateSnapshotSettingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotSettingReq struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotSettingReq", string(data)}, " ")
}
