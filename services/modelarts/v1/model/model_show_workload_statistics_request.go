package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadStatisticsRequest Request Object
type ShowWorkloadStatisticsRequest struct {

	// **参数解释**：作业所属的资源池。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`
}

func (o ShowWorkloadStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkloadStatisticsRequest", string(data)}, " ")
}
