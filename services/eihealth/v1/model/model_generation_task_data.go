package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerationTaskData 分子生成任务的请求体
type GenerationTaskData struct {

	// 生成分子数量
	NumTrials *int32 `json:"num_trials,omitempty"`

	// 强约束集合
	StrongConstraints *[]MoleculeConstraint `json:"strong_constraints,omitempty"`

	// 弱约束集合
	WeakConstraints *[]MoleculeConstraint `json:"weak_constraints,omitempty"`

	// 期望最大返回条目数（排序后取Top）
	NumExpected *int32 `json:"num_expected,omitempty"`

	// 初始化分子集合
	InitialDataset *[]string `json:"initial_dataset,omitempty"`

	BindingSite *BindingSite `json:"binding_site,omitempty"`

	// 用户已开启的自定义属性集合
	CustomProps *[]CustomProp `json:"custom_props,omitempty"`
}

func (o GenerationTaskData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerationTaskData struct{}"
	}

	return strings.Join([]string{"GenerationTaskData", string(data)}, " ")
}
