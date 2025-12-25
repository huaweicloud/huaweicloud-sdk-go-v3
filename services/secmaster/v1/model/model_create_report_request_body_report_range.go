package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportRequestBodyReportRange 数据范围
type CreateReportRequestBodyReportRange struct {

	// 起始时间
	Start string `json:"start"`

	// 终止时间
	End string `json:"end"`
}

func (o CreateReportRequestBodyReportRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportRequestBodyReportRange struct{}"
	}

	return strings.Join([]string{"CreateReportRequestBodyReportRange", string(data)}, " ")
}
