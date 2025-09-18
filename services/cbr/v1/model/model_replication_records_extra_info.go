package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplicationRecordsExtraInfo 复制记录额外信息
type ReplicationRecordsExtraInfo struct {

	// 复制进度
	Progress *int32 `json:"progress,omitempty"`

	// 失败错误码，成功时为空
	FailCode *string `json:"fail_code,omitempty"`

	// 错误原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 是否为自动调度复制
	AutoTrigger *bool `json:"auto_trigger,omitempty"`

	// 目标端的存储库id
	DestinationVaultId *string `json:"destination_vault_id,omitempty"`
}

func (o ReplicationRecordsExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationRecordsExtraInfo struct{}"
	}

	return strings.Join([]string{"ReplicationRecordsExtraInfo", string(data)}, " ")
}
