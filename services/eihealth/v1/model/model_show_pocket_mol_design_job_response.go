package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPocketMolDesignJobResponse Response Object
type ShowPocketMolDesignJobResponse struct {
	BasicInfo *DrugJobDto `json:"basic_info,omitempty"`

	Receptor *PocketMolDesignReceptorDto `json:"receptor,omitempty"`

	// 配体文件列表
	Ligands *[]DrugFile `json:"ligands,omitempty"`

	// 模型列表
	ModelList *[]BasicDrugModel `json:"model_list,omitempty"`

	// 分子量范围
	MolecularWeight *[]int32 `json:"molecular_weight,omitempty"`

	OptimizationMode *OptimizationMode `json:"optimization_mode,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

func (o ShowPocketMolDesignJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPocketMolDesignJobResponse struct{}"
	}

	return strings.Join([]string{"ShowPocketMolDesignJobResponse", string(data)}, " ")
}
