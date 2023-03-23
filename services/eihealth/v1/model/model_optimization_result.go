package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分子优化任务的返回结果
type OptimizationResult struct {

	// 任务名
	Name *string `json:"name,omitempty"`

	// 分子SMILES表达式
	Smiles *string `json:"smiles,omitempty"`

	// 总生成轮数
	NumRounds *int32 `json:"num_rounds,omitempty"`

	// 每轮生成数量
	NumTrialsPerRound *int32 `json:"num_trials_per_round,omitempty"`

	// 期望条目数
	NumExpected *int32 `json:"num_expected,omitempty"`

	// 强约束数量
	NumStrongConstraints *int32 `json:"num_strong_constraints,omitempty"`

	// 弱约束数量
	NumWeakConstraints *int32 `json:"num_weak_constraints,omitempty"`

	// 分子ADMET属性名列表
	PropNames *[]string `json:"prop_names,omitempty"`

	Original *PlainMoleculeItem `json:"original,omitempty"`

	// 分子优化结果条目
	Result *[]OptimizationResultItem `json:"result,omitempty"`
}

func (o OptimizationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptimizationResult struct{}"
	}

	return strings.Join([]string{"OptimizationResult", string(data)}, " ")
}
