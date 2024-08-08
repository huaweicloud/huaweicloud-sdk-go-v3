package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOptmJobReq 创建分子优化请求体
type CreateOptmJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	// 分子SMILES表达式
	Smiles *string `json:"smiles,omitempty"`

	MoleculeFile *DrugFile `json:"molecule_file,omitempty"`

	BindingSite *BindSiteDto `json:"binding_site,omitempty"`

	// 受体列表和受体是二选一的关系，受体列表优先级最高
	BindingSites *[]BindSiteDto `json:"binding_sites,omitempty"`

	// 弱约束集合
	WeakConstraints *[]WeakConstraintDto `json:"weak_constraints,omitempty"`

	// 强约束集合
	StrongConstraints *[]StrongConstraintDto `json:"strong_constraints,omitempty"`

	// 初始化采样权重，参数范围(0.5, 1)，不包含0.5和1，默认为0.6
	SamplerMixinWeight *float32 `json:"sampler_mixin_weight,omitempty"`

	// 基模型id
	BaseModelId *string `json:"base_model_id,omitempty"`

	// 模型id列表
	ModelIds *[]string `json:"model_ids,omitempty"`

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
