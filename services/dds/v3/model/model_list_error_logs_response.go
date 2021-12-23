package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListErrorLogsResponse struct {
	// 具体信息。

	ErrorLogList *[]ErrorlogResult `json:"error_log_list,omitempty"`
	// 数据库版本总记录数。

	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListErrorLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorLogsResponse struct{}"
	}

	return strings.Join([]string{"ListErrorLogsResponse", string(data)}, " ")
}
