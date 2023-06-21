package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建分子优化请求体
type CreateOptmJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	BindingSite *BindSiteDto `json:"binding_site,omitempty"`

	// 弱约束集合
	WeakConstraints *[]WeakConstraintDto `json:"weak_constraints,omitempty"`

	// 强约束集合
	StrongConstraints *[]StrongConstraintDto `json:"strong_constraints,omitempty"`

	// 生成分子数量
	NumTrials *int32 `json:"num_trials,omitempty"`
}

func (o CreateOptmJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOptmJobReq struct{}"
	}

	return strings.Join([]string{"CreateOptmJobReq", string(data)}, " ")
}
