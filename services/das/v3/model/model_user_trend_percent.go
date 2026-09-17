package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserTrendPercent 查询开销平均减少趋势
type UserTrendPercent struct {

	// 百分比小于等于10%的数量
	PercentLe10Count *int64 `json:"percent_le10_count,omitempty"`

	// 百分比大于10%小于等于50%的数量
	Percent10To50Count *int64 `json:"percent10_to50_count,omitempty"`

	// 百分比大于50%小于等于80%的数量
	Percent50To80Count *int64 `json:"percent50_to80_count,omitempty"`

	// 百分比大于80%的数量
	PercentGl80Count *int64 `json:"percent_gl80_count,omitempty"`
}

func (o UserTrendPercent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserTrendPercent struct{}"
	}

	return strings.Join([]string{"UserTrendPercent", string(data)}, " ")
}
