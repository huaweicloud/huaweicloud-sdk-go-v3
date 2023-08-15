package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SynthesisResultResultMolecules struct {

	// molecule的序号
	Id string `json:"id"`

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// molecule的smiles来源
	Source string `json:"source"`
}

func (o SynthesisResultResultMolecules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynthesisResultResultMolecules struct{}"
	}

	return strings.Join([]string{"SynthesisResultResultMolecules", string(data)}, " ")
}
