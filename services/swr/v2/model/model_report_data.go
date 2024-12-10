package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportData struct {

	// 时间点
	Date *string `json:"date,omitempty"`

	// 统计值
	Value *float64 `json:"value,omitempty"`
}

func (o ReportData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportData struct{}"
	}

	return strings.Join([]string{"ReportData", string(data)}, " ")
}
