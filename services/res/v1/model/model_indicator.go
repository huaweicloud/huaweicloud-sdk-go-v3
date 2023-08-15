package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Indicator
type Indicator struct {

	// 指标名称： - clickPVRate，点击PV率 - clickUVRate，点击UV率 - customize，自定义
	IndicatorName *string `json:"indicator_name,omitempty"`

	IndicatorParams *IndicatorParam `json:"indicator_params,omitempty"`
}

func (o Indicator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Indicator struct{}"
	}

	return strings.Join([]string{"Indicator", string(data)}, " ")
}
