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
}

func (o CreatePocketMolDesignJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePocketMolDesignJobReq struct{}"
	}

	return strings.Join([]string{"CreatePocketMolDesignJobReq", string(data)}, " ")
}
