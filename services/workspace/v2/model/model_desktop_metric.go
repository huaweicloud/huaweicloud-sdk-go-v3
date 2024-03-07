package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DesktopMetric struct {

	// 桌面ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 桌面池ID(仅桌面池中的桌面存在该字段)
	ResourcePoolId *string `json:"resource_pool_id,omitempty"`

	// 桌面名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 统计信息 * `desktop_usage` -  桌面使用时长(单位:秒) * `desktop_idle_duration` -  桌面空闲时长(单位:秒)
	Metric *[]Metric `json:"metric,omitempty"`
}

func (o DesktopMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopMetric struct{}"
	}

	return strings.Join([]string{"DesktopMetric", string(data)}, " ")
}
