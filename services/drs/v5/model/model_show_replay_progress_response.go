package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplayProgressResponse Response Object
type ShowReplayProgressResponse struct {

	// 回放进度的百分数
	Progress *int64 `json:"progress,omitempty"`

	// 需要解析的总数
	ParseCount *int64 `json:"parse_count,omitempty"`

	// 回放的总数
	ReplayCount *int64 `json:"replay_count,omitempty"`

	// 迁移模式
	TaskMode *string `json:"task_mode,omitempty"`

	// 迁移时间
	ProcessTime *int64 `json:"process_time,omitempty"`

	// 迁移状态
	TransferStatus *string `json:"transfer_status,omitempty"`

	// 回放最大时间
	MaxTime *int64 `json:"max_time,omitempty"`

	// 回放最小时间
	MinTime *int64 `json:"min_time,omitempty"`

	// 回放当前时间
	NowTime *int64 `json:"now_time,omitempty"`

	// 回放报告最小时间
	MinExportTime *int64 `json:"min_export_time,omitempty"`

	// 回放报告最大时间
	MaxExportTime *int64 `json:"max_export_time,omitempty"`

	// 正在回放的sql列表
	ReplaySqlNowList *[]ReplaySqlNowInfo `json:"replay_sql_now_list,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ShowReplayProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowReplayProgressResponse", string(data)}, " ")
}
