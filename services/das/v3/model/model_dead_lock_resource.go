package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeadLockResource 死锁资源信息
type DeadLockResource struct {

	// 死锁标签（keylock、objectlock、ridlock、pagelock、compilelock）
	LockLabel *string `json:"lock_label,omitempty"`

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 索引名（仅keylock展示）
	IndexName *string `json:"index_name,omitempty"`

	// 关联对象ID
	AssociatedObjectId *string `json:"associated_object_id,omitempty"`

	// 对象名称，死锁名称
	ObjectName *string `json:"object_name,omitempty"`

	// 锁模式
	LockMode *string `json:"lock_mode,omitempty"`

	// 持有者列表
	OwnerList *[]DeadLockObject `json:"owner_list,omitempty"`

	// 等待者列表
	WaiterList *[]DeadLockObject `json:"waiter_list,omitempty"`
}

func (o DeadLockResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadLockResource struct{}"
	}

	return strings.Join([]string{"DeadLockResource", string(data)}, " ")
}
