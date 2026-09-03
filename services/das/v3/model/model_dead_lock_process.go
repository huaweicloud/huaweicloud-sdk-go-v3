package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeadLockProcess 死锁会话信息
type DeadLockProcess struct {

	// 服务进程ID
	Spid *string `json:"spid,omitempty"`

	// 会话ID
	ProcessId *string `json:"process_id,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 用户名称
	LoginName *string `json:"login_name,omitempty"`

	// 任务使用的日志空间
	LogUsed *int64 `json:"log_used,omitempty"`

	// SQL语句
	Sql *string `json:"sql,omitempty"`
}

func (o DeadLockProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadLockProcess struct{}"
	}

	return strings.Join([]string{"DeadLockProcess", string(data)}, " ")
}
