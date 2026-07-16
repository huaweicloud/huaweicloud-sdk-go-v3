package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmResponseAdvancedConfig 算法高级策略：  - auto_search
type AlgorithmResponseAdvancedConfig struct {
	AutoSearch *AlgorithmResponseAdvancedConfigAutoSearch `json:"auto_search,omitempty"`
}

func (o AlgorithmResponseAdvancedConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponseAdvancedConfig struct{}"
	}

	return strings.Join([]string{"AlgorithmResponseAdvancedConfig", string(data)}, " ")
}
