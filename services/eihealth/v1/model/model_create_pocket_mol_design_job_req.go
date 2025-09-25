package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePocketMolDesignJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	Receptor *PocketMolDesignReceptorDto `json:"receptor"`

	// 配体文件列表，最多支持1个
	Ligands *[]PocketFragment `json:"ligands,omitempty"`

	// 生成分子数量
	NumTrials *int32 `json:"num_trials,omitempty"`

	// 模型id列表
	ModelIds *[]string `json:"model_ids,omitempty"`

	// 分子量范围
	MolecularWeight *[]int32 `json:"molecular_weight,omitempty"`

	OptimizationMode *OptimizationMode `json:"optimization_mode,omitempty"`

	// **参数解释**： 基模型ID。 **约束限制**： 当optimization_mode为generation时，可选择PanguDrug3D、Lingo3DMol，当为其他方式时，仅可选择可选择PanguDrug3D。 **取值范围**： - PanguDrug3D - Lingo3DMol **默认取值**： 不涉及
	BaseModelId *string `json:"base_model_id,omitempty"`
}

func (o CreatePocketMolDesignJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePocketMolDesignJobReq struct{}"
	}

	return strings.Join([]string{"CreatePocketMolDesignJobReq", string(data)}, " ")
}
