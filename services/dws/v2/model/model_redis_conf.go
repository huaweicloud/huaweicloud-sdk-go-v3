package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisConf 重分布配置信息
type RedisConf struct {

	// 重分布模式
	RedisMode string `json:"redis_mode"`

	ScheduleConf *ScheduleConf `json:"schedule_conf,omitempty"`

	// 并行作业数量
	ParallelJobs int32 `json:"parallel_jobs"`

	// 并行作业数量
	ParallelJob int32 `json:"parallel_job"`
}

func (o RedisConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConf struct{}"
	}

	return strings.Join([]string{"RedisConf", string(data)}, " ")
}
