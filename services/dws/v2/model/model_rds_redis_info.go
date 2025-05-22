package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RdsRedisInfo 重分布信息
type RdsRedisInfo struct {

	// **参数解释**： 重分布信息ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 状态。 **取值范围**： PREPARE|RUNNING|WAITING|PAUSE|PAUSING|SUCCESS|FAIL。
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
