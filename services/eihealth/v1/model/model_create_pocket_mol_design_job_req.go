package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePocketMolDesignJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	Receptor *PocketMolDesignReceptorDto `json:"receptor"`

	// 配体文件列表，最多支持1个
	Ligands *[]DrugFile `json:"ligands,omitempty"`

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
