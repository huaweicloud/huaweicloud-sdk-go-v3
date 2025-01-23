package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiSetMetricCodeThresholdReq struct {

	// 指标码
	MetricCode string `json:"metric_code"`

	// 数据库类型
	DatastoreType string `json:"datastore_type"`

	// 新阈值
	NewThreshold float64 `json:"new_threshold"`
}

func (o ApiSetMetricCodeThresholdReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiSetMetricCodeThresholdReq struct{}"
	}

	return strings.Join([]string{"ApiSetMetricCodeThresholdReq", string(data)}, " ")
}
