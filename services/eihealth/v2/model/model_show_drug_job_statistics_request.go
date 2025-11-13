package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDrugJobStatisticsRequest Request Object
type ShowDrugJobStatisticsRequest struct {

	// **参数解释**： 作业运行状态。 **约束限制**： 不涉及 **取值范围**： * WAITING：等待中 * RUNNING：运行中 * FINISHED：完成 * FAILED：失败 * CANCELLED：取消 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`
}

func (o ShowDrugJobStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrugJobStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowDrugJobStatisticsRequest", string(data)}, " ")
}
