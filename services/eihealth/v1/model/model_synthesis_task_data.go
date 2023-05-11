package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分子合成路径规划任务的请求体
type SynthesisTaskData struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 期望最大返回条目数（排序后取TopN）
	TopN int32 `json:"top_n"`

	// 预测路径的最大深度
	MaxSearchDepth int32 `json:"max_search_depth"`

	// 每个产物的最大反应数量
	MaxPredictionPerProduct int32 `json:"max_prediction_per_product"`
}

func (o SynthesisTaskData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynthesisTaskData struct{}"
	}

	return strings.Join([]string{"SynthesisTaskData", string(data)}, " ")
}
