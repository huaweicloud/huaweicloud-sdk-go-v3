package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewSessionForKillProcessTaskNewResponse Response Object
type PreviewSessionForKillProcessTaskNewResponse struct {

	// SQL限流信息列表
	Processes *[]ProcessSessionInfo `json:"processes,omitempty"`

	// 数据同步的时间
	DataSyncTime *int64 `json:"data_sync_time,omitempty"`

	// 会话总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o PreviewSessionForKillProcessTaskNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewSessionForKillProcessTaskNewResponse struct{}"
	}

	return strings.Join([]string{"PreviewSessionForKillProcessTaskNewResponse", string(data)}, " ")
}
