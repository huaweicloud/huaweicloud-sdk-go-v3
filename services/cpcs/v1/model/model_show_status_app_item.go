package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStatusAppItem struct {

	// 成功率
	SuccessRate *float64 `json:"success_rate,omitempty"`

	// 失败率
	FailRate *float64 `json:"fail_rate,omitempty"`

	// 成功次数
	SuccessCount *int64 `json:"success_count,omitempty"`

	// 总次数
	TotalCount *int64 `json:"total_count,omitempty"`

	// 平均值
	AverageValue *float64 `json:"average_value,omitempty"`

	// 总数
	TotalValue *float64 `json:"total_value,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 应用状态
	Status *string `json:"status,omitempty"`
}

func (o ShowStatusAppItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusAppItem struct{}"
	}

	return strings.Join([]string{"ShowStatusAppItem", string(data)}, " ")
}
