package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugLigandInteraction2dSvgReq 生成相互作用2D图请求体
type CreateDrugLigandInteraction2dSvgReq struct {
	ReceptorFile *ReceptorDrugFileReq `json:"receptor_file"`

	LigandFile *DrugFile `json:"ligand_file,omitempty"`

	// 小分子名称
	Name *string `json:"name,omitempty"`
}

func (o CreateDrugLigandInteraction2dSvgReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandInteraction2dSvgReq struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandInteraction2dSvgReq", string(data)}, " ")
}
