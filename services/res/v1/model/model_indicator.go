package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Indicator struct {

	// 指标名称： - clickPVRate，点击PV率 - clickUVRate，点击UV率 - customize，自定义
	IndicatorName *string `json:"indicator_name,omitempty" xml:"indicator_name"`

	IndicatorParams *IndicatorParam `json:"indicator_params,omitempty" xml:"indicator_params"`
}

func (o Indicator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Indicator struct{}"
	}

	return strings.Join([]string{"Indicator", string(data)}, " ")
}
