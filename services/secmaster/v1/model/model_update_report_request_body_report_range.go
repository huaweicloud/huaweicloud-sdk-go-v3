package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReportRequestBodyReportRange 数据范围
type UpdateReportRequestBodyReportRange struct {

	// 起始时间
	Start string `json:"start"`

	// 终止时间
	End string `json:"end"`
}

func (o UpdateReportRequestBodyReportRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportRequestBodyReportRange struct{}"
	}

	return strings.Join([]string{"UpdateReportRequestBodyReportRange", string(data)}, " ")
}
