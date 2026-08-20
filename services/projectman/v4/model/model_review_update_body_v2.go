package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReviewUpdateBodyV2 struct {

	// 评审单更新前状态。 0~32个字符。
	OldStatus *string `json:"old_status,omitempty"`

	// 评审单目标流转状态。 0~32个字符。
	Status *string `json:"status,omitempty"`

	// 计划完成时间，unix时间戳，单位：毫秒，示例：\"1759420799999\"。
	PlanEndDate *string `json:"plan_end_date,omitempty"`
}

func (o ReviewUpdateBodyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewUpdateBodyV2 struct{}"
	}

	return strings.Join([]string{"ReviewUpdateBodyV2", string(data)}, " ")
}
