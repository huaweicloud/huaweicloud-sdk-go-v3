package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountInjectionStatisticsResponse Response Object
type CountInjectionStatisticsResponse struct {

	// 生成时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// 注入列表
	InjectionStatistics *[]InjectionStatisticsBean `json:"injection_statistics,omitempty"`

	// 状态  - FINISHED：已完成  - RUNNING：运行中
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountInjectionStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInjectionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"CountInjectionStatisticsResponse", string(data)}, " ")
}
