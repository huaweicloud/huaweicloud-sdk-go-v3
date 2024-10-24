package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMetricsConfigInput 更新AOM监控配置
type UpdateMetricsConfigInput struct {

	// 是否开启AOM监控采集。
	Enable bool `json:"enable"`
}

func (o UpdateMetricsConfigInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMetricsConfigInput struct{}"
	}

	return strings.Join([]string{"UpdateMetricsConfigInput", string(data)}, " ")
}
