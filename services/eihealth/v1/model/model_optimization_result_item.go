package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OptimizationResultItem 分子优化结果条目
type OptimizationResultItem struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 分子ADMET属性值列表
	Props []PropertyValue `json:"props"`

	// 分子与初始分子的相似度
	Similarity float32 `json:"similarity"`

	// 分子所满足的弱约束数量
	NumFulfilledWeakConstraints int32 `json:"num_fulfilled_weak_constraints"`

	// 分子的打分
	Score float32 `json:"score"`
}

func (o OptimizationResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptimizationResultItem struct{}"
	}

	return strings.Join([]string{"OptimizationResultItem", string(data)}, " ")
}
