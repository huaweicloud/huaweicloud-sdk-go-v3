package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricInfo 指标信息
type MetricInfo struct {

	// 自定义指标命名空间.格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，service.item总长度最短为3，最大为32，service不能为“SYS”。namespace不能为AGT.ECS、SERVICE.BMS
	Namespace *string `json:"namespace,omitempty"`

	// 自定义指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 指标维度列表
	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`
}

func (o MetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfo struct{}"
	}

	return strings.Join([]string{"MetricInfo", string(data)}, " ")
}
