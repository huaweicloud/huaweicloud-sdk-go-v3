package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInnodbLocksResponse struct {

	// 当前持有或等待锁的事务信息
	InnodbTrx *[]InnodbTrx `json:"innodb_trx,omitempty" xml:"innodb_trx"`

	// 每个事务请求的锁以及阻塞该请求的锁的对应关系
	InnodbLockWaits *[]InnodbLockWaits `json:"innodb_lock_waits,omitempty" xml:"innodb_lock_waits"`

	// 当前持有或等待锁的事务数量
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInnodbLocksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInnodbLocksResponse struct{}"
	}

	return strings.Join([]string{"ListInnodbLocksResponse", string(data)}, " ")
}
