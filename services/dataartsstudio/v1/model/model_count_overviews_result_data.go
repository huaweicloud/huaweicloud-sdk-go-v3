package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountOverviewsResultData data，统一的返回结果的最外层数据结构。
type CountOverviewsResultData struct {
	Value *StatisticInfo `json:"value,omitempty"`
}

func (o CountOverviewsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountOverviewsResultData struct{}"
	}

	return strings.Join([]string{"CountOverviewsResultData", string(data)}, " ")
}
