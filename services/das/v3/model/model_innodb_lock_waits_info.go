package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InnodbLockWaitsInfo InnoDB锁等待信息
type InnodbLockWaitsInfo struct {

	// 请求事务的ID
	RequestingTrxId *string `json:"requesting_trx_id,omitempty"`

	// 等待的锁ID
	RequestedLockId *string `json:"requested_lock_id,omitempty"`

	// 阻塞事务的ID
	BlockingTrxId *string `json:"blocking_trx_id,omitempty"`

	// 持有的锁ID
	BlockingLockId *string `json:"blocking_lock_id,omitempty"`
}

func (o InnodbLockWaitsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InnodbLockWaitsInfo struct{}"
	}

	return strings.Join([]string{"InnodbLockWaitsInfo", string(data)}, " ")
}
