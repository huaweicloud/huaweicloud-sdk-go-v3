package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountOperationStatisticsResponse Response Object
type CountOperationStatisticsResponse struct {

	// 生成时间
	GenerateTime *string `json:"generate_time,omitempty"`

	// 风险操作统计
	OperationStatistics *[]OperationStatisticsBean `json:"operation_statistics,omitempty"`

	// 状态  - FINISHED：已完成  - RUNNING：运行中
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountOperationStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountOperationStatisticsResponse struct{}"
	}

	return strings.Join([]string{"CountOperationStatisticsResponse", string(data)}, " ")
}
