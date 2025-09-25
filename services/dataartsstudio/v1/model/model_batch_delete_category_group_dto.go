package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteCategoryGroupDto struct {

	// 待删除的数据分类id列表。
	CategoryIds []string `json:"category_ids"`
}

func (o BatchDeleteCategoryGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCategoryGroupDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteCategoryGroupDto", string(data)}, " ")
}
