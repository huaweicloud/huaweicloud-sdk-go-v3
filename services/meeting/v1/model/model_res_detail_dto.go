package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResDetailDto struct {

	// 总数。
	SumCount *int32 `json:"sumCount,omitempty"`

	// 赠送数量。
	TrialCount *int32 `json:"trialCount,omitempty"`

	// 到期数量。
	ExpiredCount *int32 `json:"expiredCount,omitempty"`

	// 即将到期数量，到期时间在30天内。
	ExpiringCount *int32 `json:"expiringCount,omitempty"`

	// 已使用数（录播存储空间、会议并发、推流并发方数暂无法查询）。
	UsedCount *int32 `json:"usedCount,omitempty"`
}

func (o ResDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResDetailDto struct{}"
	}

	return strings.Join([]string{"ResDetailDto", string(data)}, " ")
}
