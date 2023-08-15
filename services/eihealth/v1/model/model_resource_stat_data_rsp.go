package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceStatDataRsp 资源统计数据返回体
type ResourceStatDataRsp struct {

	// 统计值
	Statistic *string `json:"statistic,omitempty"`

	// 数据单位
	Unit *string `json:"unit,omitempty"`

	// 监控指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 监控资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 显卡id
	DeviceId *string `json:"device_id,omitempty"`
}

func (o ResourceStatDataRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceStatDataRsp struct{}"
	}

	return strings.Join([]string{"ResourceStatDataRsp", string(data)}, " ")
}
