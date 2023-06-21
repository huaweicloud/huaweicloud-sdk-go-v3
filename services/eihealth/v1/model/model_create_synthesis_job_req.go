package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分子合成路径规划作业请求体
type CreateSynthesisJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	Params *SynthesisParamDto `json:"params"`
}

func (o CreateSynthesisJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSynthesisJobReq struct{}"
	}

	return strings.Join([]string{"CreateSynthesisJobReq", string(data)}, " ")
}
