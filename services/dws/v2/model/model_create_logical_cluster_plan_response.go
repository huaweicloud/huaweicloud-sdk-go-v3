package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogicalClusterPlanResponse Response Object
type CreateLogicalClusterPlanResponse struct {

	// **参数解释**： 逻辑集群增删计划ID。 **取值范围**： 不涉及。
	PlanId         *string `json:"plan_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogicalClusterPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogicalClusterPlanResponse struct{}"
	}

	return strings.Join([]string{"CreateLogicalClusterPlanResponse", string(data)}, " ")
}
