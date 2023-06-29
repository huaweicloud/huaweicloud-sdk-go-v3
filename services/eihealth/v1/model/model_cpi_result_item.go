package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CpiResultItem CPI结果条目
type CpiResultItem struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 分子ADMET属性值列表
	Props []PropertyValue `json:"props"`

	// 分子与蛋白质的打分
	Score float32 `json:"score"`
}

func (o CpiResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CpiResultItem struct{}"
	}

	return strings.Join([]string{"CpiResultItem", string(data)}, " ")
}
