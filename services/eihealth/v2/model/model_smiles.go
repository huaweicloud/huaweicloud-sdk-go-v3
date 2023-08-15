package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Smiles 分子SMILES表达式
type Smiles struct {
}

func (o Smiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Smiles struct{}"
	}

	return strings.Join([]string{"Smiles", string(data)}, " ")
}
