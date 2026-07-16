package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobAlgorithmResponsePoliciesAutoSearch 超参搜索配置。
type JobAlgorithmResponsePoliciesAutoSearch struct {

	// 需要排除的超参组合。
	SkipSearchParams *string `json:"skip_search_params,omitempty"`

	// 搜索指标列表。
	RewardAttrs *[]JobAlgorithmResponsePoliciesAutoSearchRewardAttrs `json:"reward_attrs,omitempty"`

	// 搜索参数。
	SearchParams *[]JobAlgorithmResponsePoliciesAutoSearchSearchParams `json:"search_params,omitempty"`

	// 搜索算法配置。
	AlgoConfigs *[]JobAlgorithmResponsePoliciesAutoSearchAlgoConfigs `json:"algo_configs,omitempty"`
}

func (o JobAlgorithmResponsePoliciesAutoSearch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlgorithmResponsePoliciesAutoSearch struct{}"
	}

	return strings.Join([]string{"JobAlgorithmResponsePoliciesAutoSearch", string(data)}, " ")
}
