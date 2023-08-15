package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLigandSdfReq 生成sdf请求体
type CreateLigandSdfReq struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`
}

func (o CreateLigandSdfReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLigandSdfReq struct{}"
	}

	return strings.Join([]string{"CreateLigandSdfReq", string(data)}, " ")
}
