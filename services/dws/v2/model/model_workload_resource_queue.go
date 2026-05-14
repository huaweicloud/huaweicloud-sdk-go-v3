package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkloadResourceQueue struct {

	// **参数解释**： 工作负载队列短查询加速开关。 **约束限制**： 不涉及。 **取值范围**： - on：开启 - off：关闭 **默认取值**： on
	ShortQueryOptimize *string `json:"short_query_optimize,omitempty"`

	// **参数解释**： 工作负载队列短查询并发数。 **约束限制**： 不涉及。 **取值范围**： -1以上，-1表示不限制。 **默认取值**： -1
	ShortQueryConcurrencyNum *string `json:"short_query_concurrency_num,omitempty"`
}

func (o WorkloadResourceQueue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadResourceQueue struct{}"
	}

	return strings.Join([]string{"WorkloadResourceQueue", string(data)}, " ")
}
