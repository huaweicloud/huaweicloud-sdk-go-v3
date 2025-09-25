package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStarRocksSlowLogDetail struct {

	// **参数解释**：  节点ID。   **取值范围**：  只能由英文字母、数字组成，前面为UUID，后缀为no07，长度为36个字符。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**：  执行次数。   **取值范围**：   不涉及。
	Count *string `json:"count,omitempty"`

	// **参数解释**：  执行时间。   **取值范围**：   不涉及。
	Time *string `json:"time,omitempty"`

	// **参数解释**：  等待锁时间。   **取值范围**：   不涉及。
	LockTime *string `json:"lock_time,omitempty"`

	// **参数解释**：  结果行数量。   **取值范围**：   不涉及。
	RowsSent *string `json:"rows_sent,omitempty"`

	// **参数解释**：  扫描的行数量。   **取值范围**：   不涉及。
	RowsExamined *string `json:"rows_examined,omitempty"`

	// **参数解释**：   所属数据库。   **取值范围**：   gaussdb-mysql。
	Database *string `json:"database,omitempty"`

	// **参数解释**：  账号名称。   **取值范围**：   不涉及。
	Users *string `json:"users,omitempty"`

	// **参数解释**：  执行语法。   **取值范围**：   不涉及。
	QuerySample *string `json:"query_sample,omitempty"`

	// **参数解释**：  语句类型。   **取值范围**：  - INSERT - UPDATE - SELECT - DELETE - ALTER - DROP - CREATE。
	Type *string `json:"type,omitempty"`

	// **参数解释**：  发生时间，UTC时间。   **取值范围**：   不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：  IP地址。   **取值范围**：   不涉及。
	ClientIp *string `json:"client_ip,omitempty"`

	// **参数解释**：  日志单行序列号。   **取值范围**：   不涉及。
	LineNum *string `json:"line_num,omitempty"`

	// **参数解释**：  慢日志日期。   **取值范围**：   不涉及。
	SlowLogDate *int64 `json:"slow_log_date,omitempty"`
}

func (o ShowStarRocksSlowLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStarRocksSlowLogDetail struct{}"
	}

	return strings.Join([]string{"ShowStarRocksSlowLogDetail", string(data)}, " ")
}
