package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchResult 分子搜索任务的返回结果
type SearchResult struct {

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 搜索使用到的数据库集合
	Databases []string `json:"databases"`

	// 期望返回的条目数
	TopN int32 `json:"top_n"`

	// 分子ADMET属性名列表
	PropNames []string `json:"prop_names"`

	Query *PlainMoleculeItem `json:"query"`

	// 查询结果列表
	Result []SearchResultItem `json:"result"`
}

func (o SearchResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResult struct{}"
	}

	return strings.Join([]string{"SearchResult", string(data)}, " ")
}
