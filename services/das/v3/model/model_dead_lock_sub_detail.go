package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeadLockSubDetail 死锁子明细
type DeadLockSubDetail struct {

	// 事务开启时间
	LastTranStarted *string `json:"last_tran_started,omitempty"`

	// 服务进程ID
	Spid *string `json:"spid,omitempty"`

	// 该会话是否已被终止
	IsVictim *bool `json:"is_victim,omitempty"`

	// 任务使用的日志空间
	LogUsed *int64 `json:"log_used,omitempty"`

	// 锁模式（S,X,U）
	LockMode *string `json:"lock_mode,omitempty"`

	// 等待中的资源详情
	WaitResourceDesc *string `json:"wait_resource_desc,omitempty"`

	// 被锁住的对象
	ObjectOwned *string `json:"object_owned,omitempty"`

	// 请求加锁的对象
	ObjectRequested *string `json:"object_requested,omitempty"`

	// 等待资源名称
	WaitResource *string `json:"wait_resource,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 状态
	LoginName *string `json:"login_name,omitempty"`

	// 等待中的资源详情
	Status *string `json:"status,omitempty"`

	// 客户端
	ClientApp *string `json:"client_app,omitempty"`

	// SQL
	Sql *string `json:"sql,omitempty"`

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`
}

func (o DeadLockSubDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadLockSubDetail struct{}"
	}

	return strings.Join([]string{"DeadLockSubDetail", string(data)}, " ")
}
