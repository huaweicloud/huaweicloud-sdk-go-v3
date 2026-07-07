package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockAnalysisResultRequest Request Object
type ShowDeadLockAnalysisResultRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 死锁唯一标识
	DeadLockId string `json:"dead_lock_id"`

	// 死锁分析任务唯一标识
	JobId *string `json:"job_id,omitempty"`

	// 事务ID
	TransactionId *string `json:"transaction_id,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 单次返回数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDeadLockAnalysisResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockAnalysisResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDeadLockAnalysisResultRequest", string(data)}, " ")
}
