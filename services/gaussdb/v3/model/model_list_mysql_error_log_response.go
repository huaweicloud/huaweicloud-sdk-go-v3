package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListMysqlErrorLogResponse struct {
	// 错误日志具体信息。

	ErrorLogList *[]MysqlErrorLogList `json:"error_log_list,omitempty"`
	// 总记录数。

	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMysqlErrorLogResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMysqlErrorLogResponse struct{}"
	}

	return strings.Join([]string{"ListMysqlErrorLogResponse", string(data)}, " ")
}
