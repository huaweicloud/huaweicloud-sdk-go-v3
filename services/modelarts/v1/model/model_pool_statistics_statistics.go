package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolStatisticsStatistics **参数描述**： 资源池统计信息。
type PoolStatisticsStatistics struct {
	Status *PoolStatisticsStatisticsStatus `json:"status,omitempty"`
}

func (o PoolStatisticsStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatisticsStatistics struct{}"
	}

	return strings.Join([]string{"PoolStatisticsStatistics", string(data)}, " ")
}
