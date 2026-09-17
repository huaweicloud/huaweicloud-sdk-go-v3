package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKillProcessHistoryResponse Response Object
type ListKillProcessHistoryResponse struct {

	// kill会话历史记录列表
	KilledSessions *[]KillProcessHistoryInfo `json:"killed_sessions,omitempty"`

	// 历史记录总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListKillProcessHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKillProcessHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListKillProcessHistoryResponse", string(data)}, " ")
}
