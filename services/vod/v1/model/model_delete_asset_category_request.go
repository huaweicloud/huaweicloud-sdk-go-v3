package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAssetCategoryRequest struct {
	// 视频分类ID

	Id int32 `json:"id"`
}

func (o DeleteAssetCategoryRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAssetCategoryRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetCategoryRequest", string(data)}, " ")
}
