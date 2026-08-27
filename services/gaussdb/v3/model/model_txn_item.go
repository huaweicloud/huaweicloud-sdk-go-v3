package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TxnItem **参数解释**： 单个事务的进度信息。
type TxnItem struct {

	// **参数解释**： 事务的唯一标识。 **取值范围**： 不涉及。
	TrxId string `json:"trx_id"`

	// **参数解释**： 预计剩余完成时间（秒）。 **取值范围**： ＞0。
	EstimatedRemainingTime int64 `json:"estimated_remaining_time"`

	// **参数解释**： 用户会话线程ID。 **取值范围**： 不涉及。
	TrxMysqlThreadId int64 `json:"trx_mysql_thread_id"`

	// **参数解释**： 额外信息，通常是正在执行的语句。 **取值范围**： 不涉及。
	TrxQuery string `json:"trx_query"`

	// **参数解释**： 事务开始时间。 **取值范围**： 不涉及。
	TrxStarted string `json:"trx_started"`

	// **参数解释**： 事务修改的行数。 **取值范围**： ≥0。
	TrxRowsModified int64 `json:"trx_rows_modified"`
}

func (o TxnItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TxnItem struct{}"
	}

	return strings.Join([]string{"TxnItem", string(data)}, " ")
}
