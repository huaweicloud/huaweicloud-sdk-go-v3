package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InnodbLockWaits struct {

	// 申请锁资源的事务ID
	RequestingTrxId string `json:"requesting_trx_id"`

	// 申请的锁的ID
	RequestedLockId string `json:"requested_lock_id"`

	// 阻塞的事务ID
	BlockingTrxId string `json:"blocking_trx_id"`

	// 阻塞的锁的ID
	BlockingLockId string `json:"blocking_lock_id"`
}

func (o InnodbLockWaits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InnodbLockWaits struct{}"
	}

	return strings.Join([]string{"InnodbLockWaits", string(data)}, " ")
}
