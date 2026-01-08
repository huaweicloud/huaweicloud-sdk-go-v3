package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryCategoryInfoRsp struct {

	// 查询的分类ID，-1表示默认的其他分类
	Id *int32 `json:"id,omitempty"`

	// 查询的分类名称
	Name *string `json:"name,omitempty"`

	// 子层分类信息
	SubCategories *[]CategoryInfo `json:"sub_categories,omitempty"`
}

func (o QueryCategoryInfoRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCategoryInfoRsp struct{}"
	}

	return strings.Join([]string{"QueryCategoryInfoRsp", string(data)}, " ")
}
