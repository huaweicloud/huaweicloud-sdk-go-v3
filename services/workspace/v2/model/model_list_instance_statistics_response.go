package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceStatisticsResponse Response Object
type ListInstanceStatisticsResponse struct {

	// 桌面总数
	TotalCount *int64 `json:"total_count,omitempty"`

	// 未配置模型桌面数
	UnconfiguredModelCount *int64 `json:"unconfigured_model_count,omitempty"`

	// 未配置通道桌面数
	UnconfiguredChannelCount *int64 `json:"unconfigured_channel_count,omitempty"`

	// 存在风险桌面数
	RiskCount      *int64 `json:"risk_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceStatisticsResponse", string(data)}, " ")
}
