package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeadLockObject 死锁对象信息
type DeadLockObject struct {

	// 会话ID
	ProcessId *string `json:"process_id,omitempty"`

	// 服务进程ID
	Spid *string `json:"spid,omitempty"`

	// 锁模式
	LockMode *string `json:"lock_mode,omitempty"`
}

func (o DeadLockObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadLockObject struct{}"
	}

	return strings.Join([]string{"DeadLockObject", string(data)}, " ")
}
