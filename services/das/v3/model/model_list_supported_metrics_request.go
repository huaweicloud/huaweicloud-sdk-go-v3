package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportedMetricsRequest Request Object
type ListSupportedMetricsRequest struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 数据库引擎模式
	InstanceMode *string `json:"instance_mode,omitempty"`
}

func (o ListSupportedMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListSupportedMetricsRequest", string(data)}, " ")
}
