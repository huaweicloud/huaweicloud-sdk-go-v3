package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryOnlineInfoResponse Response Object
type ListHistoryOnlineInfoResponse struct {

	// 返回前端历史登录信息。查询的时间和计数之间用冒号分隔。查询的时间，按Day查询或时间段在同一天时，按小时计数，其他场景为按天计数。
	TimeCounts     *[]string `json:"time_counts,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListHistoryOnlineInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryOnlineInfoResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryOnlineInfoResponse", string(data)}, " ")
}
