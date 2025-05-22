package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisConf 重分布配置信息
type RedisConf struct {

	// **参数解释**： 重分布模式。 **取值范围**： online|offline。
	RedisMode string `json:"redis_mode"`

	ScheduleConf *ScheduleConf `json:"schedule_conf,omitempty"`

	// **参数解释**： 重分布并发数。 **取值范围**： 1~200。
	ParallelJobs int32 `json:"parallel_jobs"`

	// **参数解释**： 重分布并发数，已经废弃。 **取值范围**： 1~200。
	ParallelJob int32 `json:"parallel_job"`

	// **参数解释**： 优先级策略,支持large优先对大表进行重分布，small优先对小表进行重分布，default默认顺序进行重分布。 **取值范围**： large|small|default。
	PriorityPolicy *string `json:"priority_policy,omitempty"`

	BucketSplitInfo *BucketSplitInfo `json:"bucket_split_info,omitempty"`
}

func (o RedisConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConf struct{}"
	}

	return strings.Join([]string{"RedisConf", string(data)}, " ")
}
