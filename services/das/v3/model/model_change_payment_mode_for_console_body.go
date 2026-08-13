package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangePaymentModeForConsoleBody struct {

	// 实例ID列表
	InstanceIdList []string `json:"instance_id_list"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// true: 设置为付费, false: 设置为免费
	PaymentMode *bool `json:"payment_mode,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 是否打开全量SQL
	OpenFullSql *bool `json:"open_full_sql,omitempty"`

	// 是否打开慢SQL
	OpenSlowSql *bool `json:"open_slow_sql,omitempty"`

	// 是否打开死锁分析
	OpenDeadLock *bool `json:"open_dead_lock,omitempty"`

	// 是否打开锁阻塞
	OpenLockBlocking *bool `json:"open_lock_blocking,omitempty"`

	// 是否打开历史事务
	OpenTransaction *bool `json:"open_transaction,omitempty"`
}

func (o ChangePaymentModeForConsoleBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePaymentModeForConsoleBody struct{}"
	}

	return strings.Join([]string{"ChangePaymentModeForConsoleBody", string(data)}, " ")
}
