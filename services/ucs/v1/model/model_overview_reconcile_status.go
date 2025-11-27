package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OverviewReconcileStatus struct {

	// 配置集合总数
	ConfigSetTotalNum *int32 `json:"configSetTotalNum,omitempty"`

	// 健康状态的配置集合数量
	HealthStatusNum *int32 `json:"healthStatusNum,omitempty"`

	// 失败状态的配置集合数量
	FailedStatusNum *int32 `json:"failedStatusNum,omitempty"`

	// 未知状态的配置集合数量
	UnknownStatusNum *int32 `json:"unknownStatusNum,omitempty"`

	// 正在处理中的配置集合数量
	ProgressingStatusNum *int32 `json:"progressingStatusNum,omitempty"`
}

func (o OverviewReconcileStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverviewReconcileStatus struct{}"
	}

	return strings.Join([]string{"OverviewReconcileStatus", string(data)}, " ")
}
