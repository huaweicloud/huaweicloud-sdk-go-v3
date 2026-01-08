package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCategoryInfoRequest Request Object
type ListCategoryInfoRequest struct {

	// 视频分类ID，最多支持10个，传0表示查询所有一级分类
	Id []int32 `json:"id"`
}

func (o ListCategoryInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCategoryInfoRequest struct{}"
	}

	return strings.Join([]string{"ListCategoryInfoRequest", string(data)}, " ")
}
