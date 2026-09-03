package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FtMetricData 训练指标响应体
type FtMetricData struct {

	// 固定值 \"1.0\"，标识数据格式版本
	FormatVersion string `json:"format_version"`

	// 文件生成时间，ISO 8601 格式（如 2026-07-18T10:30:00Z）
	Timestamp *string `json:"timestamp,omitempty"`

	Metrics []FtMetric `json:"metrics"`
}

func (o FtMetricData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FtMetricData struct{}"
	}

	return strings.Join([]string{"FtMetricData", string(data)}, " ")
}
