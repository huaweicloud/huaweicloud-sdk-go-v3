package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RdsRedisInfo 重分布信息
type RdsRedisInfo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	RedisConf *RedisConf `json:"redis_conf,omitempty"`

	RedisProgress *RedisProgress `json:"redis_progress,omitempty"`

	RedisTableDetail *RedisTableDetail `json:"redis_table_detail,omitempty"`
}

func (o RdsRedisInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsRedisInfo struct{}"
	}

	return strings.Join([]string{"RdsRedisInfo", string(data)}, " ")
}
