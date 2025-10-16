package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionOptionsResult 订阅选项。
type SubscriptionOptionsResult struct {

	// 独立的分发代理。  true：使用。 false：不使用。
	IndependentAgent *bool `json:"independent_agent,omitempty"`

	// 复制架构更改。  true：可更改。 false：不可更改。
	ReplicateDdl *bool `json:"replicate_ddl,omitempty"`

	// 快照始终可用。  true：可用。 false：不可用。
	SnapshotAlwaysAvailable *bool `json:"snapshot_always_available,omitempty"`

	// 允许使用备份文件初始化。  true：允许。 false：不允许。
	AllowInitializeFromBackup *bool `json:"allow_initialize_from_backup,omitempty"`
}

func (o SubscriptionOptionsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionOptionsResult struct{}"
	}

	return strings.Join([]string{"SubscriptionOptionsResult", string(data)}, " ")
}
