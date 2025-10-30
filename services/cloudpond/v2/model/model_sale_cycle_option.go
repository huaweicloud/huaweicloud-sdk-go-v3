package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaleCycleOption 销售周期信息
type SaleCycleOption struct {
	PayMode *PayMode `json:"pay_mode,omitempty"`

	PeriodType *PeriodType `json:"period_type,omitempty"`

	// 销售周期取值
	PeriodNum *[]int32 `json:"period_num,omitempty"`
}

func (o SaleCycleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaleCycleOption struct{}"
	}

	return strings.Join([]string{"SaleCycleOption", string(data)}, " ")
}
