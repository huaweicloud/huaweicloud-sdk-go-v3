package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetaLockRequest Request Object
type ShowMetaLockRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

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

func (o ShowMetaLockRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetaLockRequest struct{}"
	}

	return strings.Join([]string{"ShowMetaLockRequest", string(data)}, " ")
}
