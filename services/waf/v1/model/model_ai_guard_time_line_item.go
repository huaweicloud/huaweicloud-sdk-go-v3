package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiGuardTimeLineItem 时间线数据项
type AiGuardTimeLineItem struct {

	// 统计指标类型： - api_detection_requests：检测请求数量 - total_matched：命中规则的请求数量
	Key string `json:"key"`

	// 时间线统计数据列表
	Timeline []TimeLinePoint `json:"timeline"`
}

func (o AiGuardTimeLineItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiGuardTimeLineItem struct{}"
	}

	return strings.Join([]string{"AiGuardTimeLineItem", string(data)}, " ")
}
