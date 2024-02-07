package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateComplexCombineReq 蛋白小分子拼接复合物接口的请求体
type GenerateComplexCombineReq struct {
	Receptor *RunReceptorPreprocessReq `json:"receptor"`

	Ligand *ReceptorDrugFileReq `json:"ligand"`
}

func (o GenerateComplexCombineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateComplexCombineReq struct{}"
	}

	return strings.Join([]string{"GenerateComplexCombineReq", string(data)}, " ")
}
