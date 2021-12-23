package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 计划周期信息
type PlanCycle struct {
	// 计划开始时间，要求用UTC时间表示。如2020-03-04

	StartDate *string `json:"start_date,omitempty"`
	// 计划结束时间，要求用UTC时间表示。如2020-03-31

	EndDate *string `json:"end_date,omitempty"`
}

func (o PlanCycle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanCycle struct{}"
	}

	return strings.Join([]string{"PlanCycle", string(data)}, " ")
}
