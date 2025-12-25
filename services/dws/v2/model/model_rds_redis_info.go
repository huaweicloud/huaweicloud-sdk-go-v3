package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RdsRedisInfo **参数解释**： 重分布信息。 **取值范围**： 不涉及。
type RdsRedisInfo struct {

	// **参数解释**： 重分布信息ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - PREPARE：准备。 - RUNNING：运行中。 - WAITING：等待中。 - PAUSE：已暂停。 - PAUSING：暂停中。 - SUCCESS：成功。 - FAIL：失败。
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
