package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListErrorLogsNewResponse struct {
	ErrorLogList *[]ErrorLog `json:"error_log_list,omitempty" xml:"error_log_list"`

	// 总记录数。
	TotalRecord    *int32 `json:"total_record,omitempty" xml:"total_record"`
	HttpStatusCode int    `json:"-"`
}

func (o ListErrorLogsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorLogsNewResponse struct{}"
	}

	return strings.Join([]string{"ListErrorLogsNewResponse", string(data)}, " ")
}
