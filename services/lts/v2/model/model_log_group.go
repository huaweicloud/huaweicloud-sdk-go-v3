package model

import (
	"encoding/json"

	"strings"
)

// 返回的日志组信息
type LogGroup struct {
	// 创建时间

	CreationTime int64 `json:"creation_time"`
	// 日志组名称

	LogGroupName string `json:"log_group_name"`
	// 日志组ID

	LogGroupId string `json:"log_group_id"`
	// 日志存储时间 天

	TtlInDays int32 `json:"ttl_in_days"`
}

func (o LogGroup) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LogGroup struct{}"
	}

	return strings.Join([]string{"LogGroup", string(data)}, " ")
}
