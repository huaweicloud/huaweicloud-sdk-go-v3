package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdmetRequest ADMET请求体
type AdmetRequest struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`
}

func (o AdmetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdmetRequest struct{}"
	}

	return strings.Join([]string{"AdmetRequest", string(data)}, " ")
}
