package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CategoryInfo struct {

	// 查询的分类ID
	Id *int32 `json:"id,omitempty"`

	// 查询的分类名称
	Name *string `json:"name,omitempty"`
}

func (o CategoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryInfo struct{}"
	}

	return strings.Join([]string{"CategoryInfo", string(data)}, " ")
}
