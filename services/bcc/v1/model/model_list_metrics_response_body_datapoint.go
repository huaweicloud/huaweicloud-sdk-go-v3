package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListMetricsResponseBodyDatapoint struct {

	// 统计维度，例如total,success,failed,
	StatisticName string `json:"statistic_name"`

	// 指标数据取值
	Value float64 `json:"value"`

	// 指标数据的单位，例如：个、GB、%
	Unit string `json:"unit"`
}

func (o ListMetricsResponseBodyDatapoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsResponseBodyDatapoint struct{}"
	}

	return strings.Join([]string{"ListMetricsResponseBodyDatapoint", string(data)}, " ")
}
