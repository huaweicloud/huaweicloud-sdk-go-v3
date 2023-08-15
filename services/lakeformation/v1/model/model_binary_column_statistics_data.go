package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BinaryColumnStatisticsData struct {

	// 列中字节数组的最大值
	MaximumLength int64 `json:"maximum_length"`

	// 列中字节数组的平均长度
	AverageLength float64 `json:"average_length"`

	// 列中空值个数
	NumberOfNull int64 `json:"number_of_null"`
}

func (o BinaryColumnStatisticsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BinaryColumnStatisticsData struct{}"
	}

	return strings.Join([]string{"BinaryColumnStatisticsData", string(data)}, " ")
}
