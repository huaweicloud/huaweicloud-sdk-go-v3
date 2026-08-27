package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEvolveTaskStatsRequest Request Object
type ListEvolveTaskStatsRequest struct {

	// **参数解释**： 关联的算法设计项目。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AlgorithmId *string `json:"algorithm_id,omitempty"`
}

func (o ListEvolveTaskStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEvolveTaskStatsRequest struct{}"
	}

	return strings.Join([]string{"ListEvolveTaskStatsRequest", string(data)}, " ")
}
