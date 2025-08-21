package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterBackupStrategyBody 开启自动创建快照策略，默认关闭。 当backupStrategy参数配置不为空时，才会开启自动创建快照策略。
type CreateClusterBackupStrategyBody struct {

	// 每天自动创建快照的时间点。只支持整点，后面需加上时区，格式为“HH:mm z”，“HH:mm”表示整点时间，“z”表示时区。比如“00:00 GMT+08:00”、“01:00 GMT+08:00”等。
	Period string `json:"period"`

	// 自动创建的快照的前缀，需要用户自己手动输入。只能包含1~32位小写字母、数字、中划线或者下划线，并且以小写字母开头。
	Prefix string `json:"prefix"`

	// 自动创建快照的保留天数。取值范围：1-90。
	Keepday int32 `json:"keepday"`

	// 快照速率参数。
	Frequency *string `json:"frequency,omitempty"`

	// 备份使用的OBS桶名称。
	Bucket *string `json:"bucket,omitempty"`

	// 快照在OBS桶中的存放路径。
	BasePath *string `json:"basePath,omitempty"`

	// 委托名称，委托给CSS，允许CSS调用您的其他云服务。   >如果bucket、basePath和agency三个参数同时为空，则系统会自动创建OBS桶和IAM代理（若创建失败，则需要手工配置正确的参数）。
	Agency *string `json:"agency,omitempty"`

	// 配置每个节点的最大备份速率（每秒），即当备份的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际备份速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxSnapshotBytesPerSeconds *string `json:"maxSnapshotBytesPerSeconds,omitempty"`

	// 配置每个节点的最大恢复速率（每秒），即当恢复的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际恢复速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxRestoreBytesPerSeconds *string `json:"maxRestoreBytesPerSeconds,omitempty"`
}

func (o CreateClusterBackupStrategyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterBackupStrategyBody struct{}"
	}

	return strings.Join([]string{"CreateClusterBackupStrategyBody", string(data)}, " ")
}
