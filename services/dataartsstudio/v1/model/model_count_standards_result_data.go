package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountStandardsResultData data，统一的返回结果的最外层数据结构。
type CountStandardsResultData struct {
	Value *StatisticStandardCoverageVo `json:"value,omitempty"`
}

func (o CountStandardsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountStandardsResultData struct{}"
	}

	return strings.Join([]string{"CountStandardsResultData", string(data)}, " ")
}
