package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResultItem 分子搜索结果条目
type SearchResultItem struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 分子所属的数据库来源
	Source string `json:"source"`

	// 分子与查询分子的相似度
	Score float32 `json:"score"`

	// 分子ADMET属性值列表
	Props []PropertyValue `json:"props"`
}

func (o SearchResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResultItem struct{}"
	}

	return strings.Join([]string{"SearchResultItem", string(data)}, " ")
}
