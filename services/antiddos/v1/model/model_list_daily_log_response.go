package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDailyLogResponse struct {
	// 弹性IP总数

	Total *int64 `json:"total,omitempty"`
	// 异常事件列表

	Logs           *[]DailyLog `json:"logs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListDailyLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDailyLogResponse struct{}"
	}

	return strings.Join([]string{"ListDailyLogResponse", string(data)}, " ")
}
