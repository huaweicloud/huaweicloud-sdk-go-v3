package model

import (
	"encoding/json"

	"strings"
)

type QueryCategoryRsp struct {
	// 分类ID。

	Id *string `json:"id,omitempty"`
	// 分类名称。

	Name *string `json:"name,omitempty"`
	// 子分类列表。

	Children *[]QueryCategoryRsp `json:"children,omitempty"`
}

func (o QueryCategoryRsp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryCategoryRsp struct{}"
	}

	return strings.Join([]string{"QueryCategoryRsp", string(data)}, " ")
}
