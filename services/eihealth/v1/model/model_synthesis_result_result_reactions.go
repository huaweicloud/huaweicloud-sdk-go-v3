package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SynthesisResultResultReactions struct {

	// 反应的序号
	Id string `json:"id"`

	// 反应物分子序号的列表
	Reactants []string `json:"reactants"`

	// 产物分子序号
	Product string `json:"product"`
}

func (o SynthesisResultResultReactions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynthesisResultResultReactions struct{}"
	}

	return strings.Join([]string{"SynthesisResultResultReactions", string(data)}, " ")
}
