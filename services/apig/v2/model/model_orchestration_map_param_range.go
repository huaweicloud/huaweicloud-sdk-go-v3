package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrchestrationMapParamRange 参数区间编排配置。
type OrchestrationMapParamRange struct {

	// 区间起始值。  为可以转换成integer的string，转换后的range_start的范围为0-9223372036854775807， range_start不大于range_end。
	RangeStart *string `json:"range_start,omitempty"`

	// 区间终止值。  为可以转换成integer的string，转换后的range_end的范围为0-9223372036854775807， range_start不大于range_end。
	RangeEnd *string `json:"range_end,omitempty"`
}

func (o OrchestrationMapParamRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrchestrationMapParamRange struct{}"
	}

	return strings.Join([]string{"OrchestrationMapParamRange", string(data)}, " ")
}
