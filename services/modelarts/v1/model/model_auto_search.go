package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoSearch 超参搜索配置。
type AutoSearch struct {

	// 需要排除的超参组合。
	SkipSearchParams *string `json:"skip_search_params,omitempty"`

	// 搜索指标列表。
	RewardAttrs *[]RewardAttrs `json:"reward_attrs,omitempty"`

	// 搜索参数。
	SearchParams *[]SearchParams `json:"search_params,omitempty"`

	// 搜索算法配置。
	AlgoConfigs *[]AlgoConfigs `json:"algo_configs,omitempty"`
}

func (o AutoSearch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoSearch struct{}"
	}

	return strings.Join([]string{"AutoSearch", string(data)}, " ")
}
