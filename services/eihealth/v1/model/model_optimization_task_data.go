package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OptimizationTaskData 分子优化任务的请求体
type OptimizationTaskData struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 生成分子数量
	NumTrials *int32 `json:"num_trials,omitempty"`

	// 强约束集合
	StrongConstraints *[]MoleculeConstraint `json:"strong_constraints,omitempty"`

	// 弱约束集合
	WeakConstraints *[]MoleculeConstraint `json:"weak_constraints,omitempty"`

	// 期望最大返回条目数（排序后取Top）
	NumExpected *int32 `json:"num_expected,omitempty"`

	BindingSite *BindingSite `json:"binding_site,omitempty"`

	// 用户已开启的自定义属性集合
	CustomProps *[]CustomProp `json:"custom_props,omitempty"`
}

func (o OptimizationTaskData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptimizationTaskData struct{}"
	}

	return strings.Join([]string{"OptimizationTaskData", string(data)}, " ")
}
