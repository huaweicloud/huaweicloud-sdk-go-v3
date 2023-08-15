package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerationResultItem 分子生成结果条目
type GenerationResultItem struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 分子ADMET属性值列表
	Props []PropertyValue `json:"props"`

	// 分子所满足的弱约束数量
	NumFulfilledWeakConstraints int32 `json:"num_fulfilled_weak_constraints"`

	// 分子的打分
	Score float32 `json:"score"`
}

func (o GenerationResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerationResultItem struct{}"
	}

	return strings.Join([]string{"GenerationResultItem", string(data)}, " ")
}
