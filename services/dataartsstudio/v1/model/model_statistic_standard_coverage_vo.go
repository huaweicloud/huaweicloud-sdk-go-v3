package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticStandardCoverageVo struct {

	// 字段总数，填写String类型替代Long类型。
	AllColNum *string `json:"all_col_num,omitempty"`

	// 关联标准字段数，填写String类型替代Long类型。
	ColNum *string `json:"col_num,omitempty"`

	// 标准覆盖率。
	Coverage *float64 `json:"coverage,omitempty"`

	// 引用表数组。
	Details *[]AllTableVo `json:"details,omitempty"`
}

func (o StatisticStandardCoverageVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticStandardCoverageVo struct{}"
	}

	return strings.Join([]string{"StatisticStandardCoverageVo", string(data)}, " ")
}
