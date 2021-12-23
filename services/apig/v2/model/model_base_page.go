package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页基本信息
type BasePage struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
}

func (o BasePage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePage struct{}"
	}

	return strings.Join([]string{"BasePage", string(data)}, " ")
}
