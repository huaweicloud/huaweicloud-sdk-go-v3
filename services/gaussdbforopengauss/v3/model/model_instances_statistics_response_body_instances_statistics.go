package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstancesStatisticsResponseBodyInstancesStatistics struct {

	// 实例状态
	Status string `json:"status"`

	// 实例数量
	Count int32 `json:"count"`
}

func (o InstancesStatisticsResponseBodyInstancesStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesStatisticsResponseBodyInstancesStatistics struct{}"
	}

	return strings.Join([]string{"InstancesStatisticsResponseBodyInstancesStatistics", string(data)}, " ")
}
