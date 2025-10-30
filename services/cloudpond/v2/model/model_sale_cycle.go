package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaleCycle 销售周期对象
type SaleCycle struct {
	PayMode *PayMode `json:"pay_mode,omitempty"`

	PeriodType *PeriodType `json:"period_type,omitempty"`

	// 周期数量
	PeriodNum *int32 `json:"period_num,omitempty"`
}

func (o SaleCycle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaleCycle struct{}"
	}

	return strings.Join([]string{"SaleCycle", string(data)}, " ")
}
