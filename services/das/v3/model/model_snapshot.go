package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Snapshot 锁快照信息
type Snapshot struct {

	// 快照ID
	Id *int64 `json:"id,omitempty"`

	// 快照状态。取值范围：0（等待中）、1（运行中）、2（失败）、3（成功）
	Status *int32 `json:"status,omitempty"`

	// 锁快照创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 是否找到有锁。取值范围：0（否）、1（是）
	FindLock *int32 `json:"find_lock,omitempty"`
}

func (o Snapshot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Snapshot struct{}"
	}

	return strings.Join([]string{"Snapshot", string(data)}, " ")
}
