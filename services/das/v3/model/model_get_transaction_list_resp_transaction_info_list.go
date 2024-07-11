package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetTransactionListRespTransactionInfoList struct {

	// 事务持续时间
	LastSec int32 `json:"last_sec"`

	// 等待锁数量
	WaitLocks int32 `json:"wait_locks"`

	// 持有锁数量
	HoldLocks int32 `json:"hold_locks"`

	// 发生时间
	OccurrenceTime int32 `json:"occurrence_time"`

	// 事务内容
	Detail string `json:"detail"`

	// 收集时间
	CollectTime int64 `json:"collect_time"`
}

func (o GetTransactionListRespTransactionInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTransactionListRespTransactionInfoList struct{}"
	}

	return strings.Join([]string{"GetTransactionListRespTransactionInfoList", string(data)}, " ")
}
