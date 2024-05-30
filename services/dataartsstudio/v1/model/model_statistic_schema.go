package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticSchema struct {

	// 本月新增。
	Increase *int32 `json:"increase,omitempty"`

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// 标准覆盖率。
	StandardCoverage *float64 `json:"standard_coverage,omitempty"`
}

func (o StatisticSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticSchema struct{}"
	}

	return strings.Join([]string{"StatisticSchema", string(data)}, " ")
}
