package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogsResponse Response Object
type ListSlowLogsResponse struct {

	// 具体信息。
	SlowLogList *[]SlowlogResult `json:"slow_log_list,omitempty"`

	// 数据库版本总记录数。
	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogsResponse", string(data)}, " ")
}
