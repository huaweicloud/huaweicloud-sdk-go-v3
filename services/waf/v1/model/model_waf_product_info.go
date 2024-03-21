package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WafProductInfo struct {

	// waf规格   -  professional：标准   - enterprise：专业   - ultimate：铂金版
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 订购周期类型 month: 月；year: 年
	PeriodType *string `json:"period_type,omitempty"`

	// 订购周期数
	PeriodNum *int32 `json:"period_num,omitempty"`
}

func (o WafProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafProductInfo struct{}"
	}

	return strings.Join([]string{"WafProductInfo", string(data)}, " ")
}
