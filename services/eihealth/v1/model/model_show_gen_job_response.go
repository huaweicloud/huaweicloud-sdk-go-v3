package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGenJobResponse Response Object
type ShowGenJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	// 分子表达式列表
	SmilesList *[]string `json:"smiles_list,omitempty"`

	MoleculeFile *DrugFile `json:"molecule_file,omitempty"`

	// 生成分子数量
	NumTrials *int32 `json:"num_trials,omitempty"`

	// 初始化数据集的分子数目。当为-1时，表示分子数目未知
	InitialDatasetSize *int32 `json:"initial_dataset_size,omitempty"`

	// 靶点列表
	BindingSites *[]BindSiteDto `json:"binding_sites,omitempty"`

	// 弱约束集合
	WeakConstraints *[]WeakConstraintDto `json:"weak_constraints,omitempty"`

	// 强约束集合
	StrongConstraints *[]StrongConstraintDto `json:"strong_constraints,omitempty"`

	BaseModel *BaseModel `json:"base_model,omitempty"`

	// 模型列表
	Models         *[]BasicDrugModel `json:"models,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowGenJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGenJobResponse struct{}"
	}

	return strings.Join([]string{"ShowGenJobResponse", string(data)}, " ")
}
