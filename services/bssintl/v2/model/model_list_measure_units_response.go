package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMeasureUnitsResponse Response Object
type ListMeasureUnitsResponse struct {

	// 度量单位信息，具体参见表2。
	MeasureUnits   *[]MeasureUnitRest `json:"measure_units,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListMeasureUnitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMeasureUnitsResponse struct{}"
	}

	return strings.Join([]string{"ListMeasureUnitsResponse", string(data)}, " ")
}
