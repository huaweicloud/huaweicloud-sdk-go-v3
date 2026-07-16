package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgoConfigs 搜索算法配置。
type AlgoConfigs struct {

	// 搜索算法名称。
	Name *string `json:"name,omitempty"`

	// 搜索算法参数。
	Params *[]AutoSearchAlgoConfigParameter `json:"params,omitempty"`
}

func (o AlgoConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgoConfigs struct{}"
	}

	return strings.Join([]string{"AlgoConfigs", string(data)}, " ")
}
