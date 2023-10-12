package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisProgress 重分布进度信息
type RedisProgress struct {

	// 已完成字节数
	BytesDone *int64 `json:"bytes_done,omitempty"`

	// 剩余字节数
	ByteLeft *int64 `json:"byte_left,omitempty"`

	// 完成表数量
	TablesDone *int32 `json:"tables_done,omitempty"`

	// 剩余表数量
	TablesLeft *int32 `json:"tables_left,omitempty"`

	// 表重分布进度
	TableProgress *int32 `json:"table_progress,omitempty"`

	// 总进度
	TotalProgress *int32 `json:"total_progress,omitempty"`

	// 重分布比例
	RedisRate *string `json:"redis_rate,omitempty"`

	// 预计时间
	EstimatedTime *string `json:"estimated_time,omitempty"`

	// 是否已完成
	Completed *bool `json:"completed,omitempty"`

	// 是否完成初始化
	Initialed *bool `json:"initialed,omitempty"`

	// 失败总数
	FailCount *int32 `json:"fail_count,omitempty"`

	// cm_ctl 结果
	Redistributing *bool `json:"redistributing,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 是否用户暂停
	PauseByUser *bool `json:"pause_by_user,omitempty"`
}

func (o RedisProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisProgress struct{}"
	}

	return strings.Join([]string{"RedisProgress", string(data)}, " ")
}
