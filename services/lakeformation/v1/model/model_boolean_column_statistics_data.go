package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BooleanColumnStatisticsData struct {

	// 列中为真的数量
	NumberOfTrue int64 `json:"number_of_true"`

	// 列中为假的数量
	NumberOfFalse int64 `json:"number_of_false"`

	// 列中为空的数量
	NumberOfNull int64 `json:"number_of_null"`
}

func (o BooleanColumnStatisticsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BooleanColumnStatisticsData struct{}"
	}

	return strings.Join([]string{"BooleanColumnStatisticsData", string(data)}, " ")
}
