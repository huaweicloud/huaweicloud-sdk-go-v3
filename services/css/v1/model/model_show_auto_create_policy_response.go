package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoCreatePolicyResponse Response Object
type ShowAutoCreatePolicyResponse struct {

	// 快照保留的天数。
	Keepday *int32 `json:"keepday,omitempty"`

	// 每天快照创建时刻。
	Period *string `json:"period,omitempty"`

	// 自动创建快照的执行频次。
	Frequency *string `json:"frequency,omitempty"`

	// 快照命名前缀，需要用户自己手动输入。
	Prefix *string `json:"prefix,omitempty"`

	// 快照存放的OBS桶的桶名。
	Bucket *string `json:"bucket,omitempty"`

	// 快照在OBS桶中的存放路径。
	BasePath *string `json:"basePath,omitempty"`

	// 访问OBS桶用到的委托。
	Agency *string `json:"agency,omitempty"`

	// 是否开启自动创建快照策略。 - true：表示开启自动创建快照策略。 - false：表示关闭自动创建快照策略。
	Enable *string `json:"enable,omitempty"`

	// 需要备份的索引名。
	Indices *string `json:"indices,omitempty"`

	// 快照ID。
	SnapshotCmkId *string `json:"snapshotCmkId,omitempty"`

	// 配置每个节点的最大备份速率（每秒），即当备份的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际备份速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxSnapshotBytesPerSeconds *string `json:"maxSnapshotBytesPerSeconds,omitempty"`

	// 配置每个节点的最大恢复速率（每秒），即当恢复的速率超过该值时会被限流，避免速率太大导致资源占用过高，影响系统稳定性。实际恢复速率不一定能达到该值，会受OBS、磁盘等影响。
	MaxRestoreBytesPerSeconds *string `json:"maxRestoreBytesPerSeconds,omitempty"`
	HttpStatusCode            int     `json:"-"`
}

func (o ShowAutoCreatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoCreatePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoCreatePolicyResponse", string(data)}, " ")
}
