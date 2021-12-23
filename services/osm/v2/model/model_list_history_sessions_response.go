package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistorySessionsResponse struct {
	// 总记录数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 会话列表

	SessionList    *[]OperateHistorySession `json:"session_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListHistorySessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistorySessionsResponse struct{}"
	}

	return strings.Join([]string{"ListHistorySessionsResponse", string(data)}, " ")
}
