package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGrowthRateResponse Response Object
type ShowGrowthRateResponse struct {

	// 环比值。
	GrowthRate     *float64 `json:"growth_rate,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowGrowthRateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGrowthRateResponse struct{}"
	}

	return strings.Join([]string{"ShowGrowthRateResponse", string(data)}, " ")
}
