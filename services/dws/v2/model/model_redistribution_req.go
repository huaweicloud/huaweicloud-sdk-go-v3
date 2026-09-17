package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedistributionReq **参数解释**： 重分布请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RedistributionReq struct {

	// **参数解释**： 重分布模式 **约束限制**： 不涉及。 **取值范围**： online：在线重分布； offline：离线重分布； **默认取值**： offline
	RedisMode string `json:"redis_mode"`

	// **参数解释**： 重分布并发数。 **约束限制**： 不涉及。 **取值范围**： 4~200 **默认取值**： 4
	ParallelJobs int32 `json:"parallel_jobs"`
}

func (o RedistributionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedistributionReq struct{}"
	}

	return strings.Join([]string{"RedistributionReq", string(data)}, " ")
}
