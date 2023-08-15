package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SynthesisParamDto 分子合成路径规划参数列表
type SynthesisParamDto struct {

	// 期望最大返回条目数（排序后取TopN）
	TopN int32 `json:"top_n"`

	// 预测路径的最大深度
	MaxSearchDepth int32 `json:"max_search_depth"`

	// 搜索最大时间，单位:分钟
	TimeLimit int32 `json:"time_limit"`

	// 每个产物的最大反应数量
	MaxPredictionPerProduct int32 `json:"max_prediction_per_product"`
}

func (o SynthesisParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynthesisParamDto struct{}"
	}

	return strings.Join([]string{"SynthesisParamDto", string(data)}, " ")
}
