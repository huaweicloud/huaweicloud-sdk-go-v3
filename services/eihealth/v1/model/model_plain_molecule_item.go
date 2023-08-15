package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PlainMoleculeItem struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 分子ADMET属性值列表
	Props []PropertyValue `json:"props"`
}

func (o PlainMoleculeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlainMoleculeItem struct{}"
	}

	return strings.Join([]string{"PlainMoleculeItem", string(data)}, " ")
}
