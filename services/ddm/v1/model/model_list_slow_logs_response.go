package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogsResponse Response Object
type ListSlowLogsResponse struct {

	// DDM慢sql日志条数。
	TotalRecord *int32 `json:"total_record,omitempty"`

	// DDM慢sql日志信息列表的集合。
	SlowLogList    *[]SlowLogs `json:"slow_log_list,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogsResponse", string(data)}, " ")
}
