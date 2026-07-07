package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDeadLockTopologyGraphRespLocks struct {

	// 锁节点唯一标识
	LockId string `json:"lock_id"`

	// 事务节点唯一标识
	TransactionId string `json:"transaction_id"`

	// 索引名字
	IndexName string `json:"index_name"`

	// 锁类型
	LockType string `json:"lock_type"`

	// 锁模式
	LockMode string `json:"lock_mode"`

	// 锁状态
	LockStatus string `json:"lock_status"`

	// 表空间ID
	SpaceId int64 `json:"space_id"`

	// 页号
	PageNo int64 `json:"page_no"`

	// 堆号
	HeapNo int64 `json:"heap_no"`

	// 操作的表名
	TableName string `json:"table_name"`

	// 是否主键索引
	PrimaryKey bool `json:"primary_key"`

	// 锁定的字段数据
	LockedData []ShowDeadLockTopologyGraphRespLockedData `json:"locked_data"`

	// 是否未知锁
	Unknown bool `json:"unknown"`
}

func (o ShowDeadLockTopologyGraphRespLocks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyGraphRespLocks struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyGraphRespLocks", string(data)}, " ")
}
