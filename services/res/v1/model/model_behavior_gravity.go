package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重力衰减因子
type BehaviorGravity struct {

	// 衰减因子。
	WeakenFactor *float64 `json:"weaken_factor,omitempty"`

	// 行为次数统计方法： - pv，访问量 - uv，独立访客
	ViewType *string `json:"view_type,omitempty"`

	// 算法类型: - normal，通用 - time，时间
	AlgoType *string `json:"algo_type,omitempty"`
}

func (o BehaviorGravity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BehaviorGravity struct{}"
	}

	return strings.Join([]string{"BehaviorGravity", string(data)}, " ")
}
