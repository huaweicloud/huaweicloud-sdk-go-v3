package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnosisInfo struct {

	// 指标名称。
	MetricName string `json:"metric_name"`

	// 实例数量。
	Count int32 `json:"count"`
}

func (o DiagnosisInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisInfo struct{}"
	}

	return strings.Join([]string{"DiagnosisInfo", string(data)}, " ")
}
