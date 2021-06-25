package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAssetCategoryRequest struct {
	// 视频分类ID

	Id int32 `json:"id"`
}

func (o ListAssetCategoryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAssetCategoryRequest struct{}"
	}

	return strings.Join([]string{"ListAssetCategoryRequest", string(data)}, " ")
}
