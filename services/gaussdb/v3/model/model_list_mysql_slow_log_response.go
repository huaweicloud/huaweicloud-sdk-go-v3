package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListMysqlSlowLogResponse struct {
	// 错误日志具体信息。

	SlowLogList *[]MysqlSlowLogList `json:"slow_log_list,omitempty"`
	// 慢日志阈值。

	LongQueryTime *string `json:"long_query_time,omitempty"`
	// 总记录数。

	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMysqlSlowLogResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMysqlSlowLogResponse struct{}"
	}

	return strings.Join([]string{"ListMysqlSlowLogResponse", string(data)}, " ")
}
