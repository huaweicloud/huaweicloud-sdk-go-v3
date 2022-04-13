package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAssetCategoryRequest struct {
	// 视频分类ID。  若设置为0，则查询所有一级分类。

	Id int32 `json:"id"`
}

func (o ListAssetCategoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetCategoryRequest struct{}"
	}

	return strings.Join([]string{"ListAssetCategoryRequest", string(data)}, " ")
}
