package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExceptionalDates 例外的日期，valid_dates表示需要执行的日期，invalid_dates表示不执行的日期。其优先级最高，优先级：ExceptionalDates > weekdays > start_time-endtime
type ExceptionalDates struct {

	// 例外日期
	ValidDates *[]string `json:"valid_dates,omitempty"`

	// 无效日期
	InvalidDates *[]string `json:"invalid_dates,omitempty"`
}

func (o ExceptionalDates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceptionalDates struct{}"
	}

	return strings.Join([]string{"ExceptionalDates", string(data)}, " ")
}
