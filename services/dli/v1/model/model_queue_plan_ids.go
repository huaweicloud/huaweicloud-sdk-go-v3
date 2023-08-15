package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueuePlanIds
type QueuePlanIds struct {

	// 表示SQL语句的类型
	PlanIds []int64 `json:"plan_ids"`
}

func (o QueuePlanIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueuePlanIds struct{}"
	}

	return strings.Join([]string{"QueuePlanIds", string(data)}, " ")
}
