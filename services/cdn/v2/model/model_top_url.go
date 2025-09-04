package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopUrl top_url配置
type TopUrl struct {

	// 配置开关
	Enable *bool `json:"enable,omitempty"`

	// 热点统计配置指标的上报数量。如top_url 100、top_url 1000
	Limit *int32 `json:"limit,omitempty"`

	// 热点统计类指标是否支持按状态码上报
	SortByCode *bool `json:"sort_by_code,omitempty"`
}

func (o TopUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrl struct{}"
	}

	return strings.Join([]string{"TopUrl", string(data)}, " ")
}
