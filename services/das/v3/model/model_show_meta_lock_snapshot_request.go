package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetaLockSnapshotRequest Request Object
type ShowMetaLockSnapshotRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 元数据锁快照ID
	Id int32 `json:"id"`

	// 线程ID
	ThreadId *string `json:"thread_id,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 锁状态
	LockStatus *string `json:"lock_status,omitempty"`

	// 锁类型
	LockType *string `json:"lock_type,omitempty"`
}

func (o ShowMetaLockSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetaLockSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowMetaLockSnapshotRequest", string(data)}, " ")
}
