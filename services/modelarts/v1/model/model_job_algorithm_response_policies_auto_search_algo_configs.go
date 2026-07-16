package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobAlgorithmResponsePoliciesAutoSearchAlgoConfigs struct {

	// 搜索算法名称。
	Name *string `json:"name,omitempty"`

	// 搜索算法参数。
	Params *[]AutoSearchAlgoConfigParameter `json:"params,omitempty"`
}

func (o JobAlgorithmResponsePoliciesAutoSearchAlgoConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlgorithmResponsePoliciesAutoSearchAlgoConfigs struct{}"
	}

	return strings.Join([]string{"JobAlgorithmResponsePoliciesAutoSearchAlgoConfigs", string(data)}, " ")
}
