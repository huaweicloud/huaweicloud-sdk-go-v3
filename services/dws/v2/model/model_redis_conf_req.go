package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisConfReq 重分布配置请求信息
type RedisConfReq struct {

	// **参数解释**： 重分布模式，不同模式对业务影响不同，修改建议联系运维人员评估后再决定。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： offLine：离线模式。 onLine：在线模式。 **默认取值**： offLine
	RedisMode string `json:"redis_mode"`

	// **参数解释**： 并行作业数量，默认4。 **约束限制**： 不涉及。 **取值范围**： 1~200 **默认取值**： 4
	ParallelJobs int32 `json:"parallel_jobs"`
}

func (o RedisConfReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConfReq struct{}"
	}

	return strings.Join([]string{"RedisConfReq", string(data)}, " ")
}
