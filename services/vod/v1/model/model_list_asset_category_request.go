package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAssetCategoryRequest struct {
	// 视频分类ID。 若设置为0，则查询所有一级分类。

	Id int32 `json:"id"`
}

func (o ListAssetCategoryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAssetCategoryRequest struct{}"
	}

	return strings.Join([]string{"ListAssetCategoryRequest", string(data)}, " ")
}
