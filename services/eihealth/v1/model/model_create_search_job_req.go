package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSearchJobReq struct {
	BasicInfo *CreateDrugJobBasicInfo `json:"basic_info"`

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 分子骨架表达式
	Scaffold *string `json:"scaffold,omitempty"`

	// 最相似的top_n个
	TopN *int32 `json:"top_n,omitempty"`

	// 可供搜索分子的公共数据库名称列表
	Databases *[]string `json:"databases,omitempty"`

	// 可供搜索分子的自定义数据库id列表
	CustomDatabases *[]string `json:"custom_databases,omitempty"`

	// 模型id列表
	ModelIds *[]string `json:"model_ids,omitempty"`

	SearchMethod *SearchType `json:"search_method,omitempty"`
}

func (o CreateSearchJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchJobReq struct{}"
	}

	return strings.Join([]string{"CreateSearchJobReq", string(data)}, " ")
}
