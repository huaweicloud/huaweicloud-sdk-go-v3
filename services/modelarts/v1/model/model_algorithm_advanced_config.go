package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmAdvancedConfig 算法高级策略：  - auto_search
type AlgorithmAdvancedConfig struct {
	AutoSearch *AutoSearch `json:"auto_search,omitempty"`
}

func (o AlgorithmAdvancedConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmAdvancedConfig struct{}"
	}

	return strings.Join([]string{"AlgorithmAdvancedConfig", string(data)}, " ")
}
