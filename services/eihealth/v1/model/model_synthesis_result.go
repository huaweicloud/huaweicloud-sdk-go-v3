package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SynthesisResult 分子合成路径规划任务的返回结果
type SynthesisResult struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 期望返回的条目数
	TopN int32 `json:"top_n"`

	// 期望搜索的最大深度
	MaxSearchDepth int32 `json:"max_search_depth"`

	// 期望每个产物的最多反应数量
	MaxPredictionPerProduct int32 `json:"max_prediction_per_product"`

	Result *SynthesisResultResult `json:"result"`
}

func (o SynthesisResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynthesisResult struct{}"
	}

	return strings.Join([]string{"SynthesisResult", string(data)}, " ")
}
