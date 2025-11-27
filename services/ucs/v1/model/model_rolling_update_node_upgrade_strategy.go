package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RollingUpdateNodeUpgradeStrategy struct {

	// 最大不可用个数
	MaxUnavailable *interface{} `json:"maxUnavailable,omitempty"`

	// 允许超出期望的最大个数
	MaxSurge *interface{} `json:"maxSurge,omitempty"`

	// 删除策略：Random、Oldest、Newest
	DeletePolicy *string `json:"deletePolicy,omitempty"`
}

func (o RollingUpdateNodeUpgradeStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollingUpdateNodeUpgradeStrategy struct{}"
	}

	return strings.Join([]string{"RollingUpdateNodeUpgradeStrategy", string(data)}, " ")
}
