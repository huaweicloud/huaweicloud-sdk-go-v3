package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedistributionConf **参数解释**： 重分布可设置的配置信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RedistributionConf struct {

	// **参数解释**： 并发数，默认并发数为4。 **约束限制**： 不涉及。 **取值范围**： 1~200 **默认取值**： 不涉及。
	ParallelJobs int32 `json:"parallel_jobs"`

	// **参数解释**： 重分布优先级策略。 **约束限制**： 不涉及。 **取值范围**： - default：默认策略 - small：小表优先 - large：大表优先  **默认取值**： 不涉及。
	PriorityPolicy string `json:"priority_policy"`
}

func (o RedistributionConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedistributionConf struct{}"
	}

	return strings.Join([]string{"RedistributionConf", string(data)}, " ")
}
