package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGenJobReq 创建分子生成请求体
type CreateGenJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	// 分子表达式列表
	SmilesList *[]string `json:"smiles_list,omitempty"`

	MoleculeFile *DrugFile `json:"molecule_file,omitempty"`

	// 靶点列表
	BindingSites *[]BindSiteDto `json:"binding_sites,omitempty"`

	// 弱约束集合
	WeakConstraints *[]WeakConstraintDto `json:"weak_constraints,omitempty"`

	// 强约束集合
	StrongConstraints *[]StrongConstraintDto `json:"strong_constraints,omitempty"`

	// 基模型id
	BaseModelId *string `json:"base_model_id,omitempty"`

	// 模型id列表
	ModelIds *[]string `json:"model_ids,omitempty"`

	// 生成分子数量
	NumTrials *int32 `json:"num_trials,omitempty"`
}

func (o CreateGenJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGenJobReq struct{}"
	}

	return strings.Join([]string{"CreateGenJobReq", string(data)}, " ")
}
