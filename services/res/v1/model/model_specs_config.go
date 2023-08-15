package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecsConfig
type SpecsConfig struct {

	// 离线计算规格。
	Offline string `json:"offline"`

	// 实时计算规格。
	Nearline *string `json:"nearline,omitempty"`

	// 深度学习计算规格。
	Rank *string `json:"rank,omitempty"`

	// 在线服务最大并发数。
	OnlineTps *int32 `json:"online_tps,omitempty"`
}

func (o SpecsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecsConfig struct{}"
	}

	return strings.Join([]string{"SpecsConfig", string(data)}, " ")
}
