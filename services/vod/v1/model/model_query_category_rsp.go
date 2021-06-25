package model

import (
	"encoding/json"

	"strings"
)

type QueryCategoryRsp struct {
	// 视频分类ID<br/>

	Id *string `json:"id,omitempty"`
	// 视频分类名称<br/>

	Name *string `json:"name,omitempty"`
	// 子分类信息<br/>

	Children *[]QueryCategoryRsp `json:"children,omitempty"`
}

func (o QueryCategoryRsp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryCategoryRsp struct{}"
	}

	return strings.Join([]string{"QueryCategoryRsp", string(data)}, " ")
}
