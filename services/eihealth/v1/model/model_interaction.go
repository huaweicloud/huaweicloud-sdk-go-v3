package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Interaction 分子与蛋白间的相互作用力
type Interaction struct {

	// 靶点，只支持target1或target2。
	BindingSite string `json:"binding_site"`

	Type *InteractionType `json:"type"`

	// 氨基酸
	AminoAcid string `json:"amino_acid"`
}

func (o Interaction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Interaction struct{}"
	}

	return strings.Join([]string{"Interaction", string(data)}, " ")
}
