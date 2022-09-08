package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Cost struct {

	// 时间维度值。 如按天维度，该值为yyyy-mm-dd如按月维度，该值为yyyy-mm
	TimeDimensionValue *string `json:"time_dimension_value,omitempty"`

	// 时间单位。 1：天2：月
	TimeMeasureId *int32 `json:"time_measure_id,omitempty"`

	// 应付或实付的成本金额或均摊金额，具体取决于请求参数。
	Amount *string `json:"amount,omitempty"`

	// 官网价金额。
	OfficialAmount *string `json:"official_amount,omitempty"`
}

func (o Cost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cost struct{}"
	}

	return strings.Join([]string{"Cost", string(data)}, " ")
}
