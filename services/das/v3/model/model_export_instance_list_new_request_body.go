package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportInstanceListNewRequestBody 导出实例列表请求体
type ExportInstanceListNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 数据库引擎类型
	EngineGroup string `json:"engine_group"`

	// 实例状态，取值范围：normal（正常）、abnormal（异常）、metricAbnormal（指标异常）、dataDiskFull（磁盘不足）、all（所有）
	InstanceStatus string `json:"instance_status"`

	// 排序条件
	OrderValue *string `json:"order_value,omitempty"`

	// 指标名称
	MetricNames *[]string `json:"metric_names,omitempty"`
}

func (o ExportInstanceListNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstanceListNewRequestBody struct{}"
	}

	return strings.Join([]string{"ExportInstanceListNewRequestBody", string(data)}, " ")
}
