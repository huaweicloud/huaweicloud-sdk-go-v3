package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分子SMILES表达式
type Smiles struct {
}

func (o Smiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Smiles struct{}"
	}

	return strings.Join([]string{"Smiles", string(data)}, " ")
}
