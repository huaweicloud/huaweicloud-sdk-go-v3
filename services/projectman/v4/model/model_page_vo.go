package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageVo 分页参数返回体
type PageVo struct {

	// 当前页
	Page *string `json:"page,omitempty"`

	// 每页条数
	Size *string `json:"size,omitempty"`

	// 数据总数
	Count *string `json:"count,omitempty"`
}

func (o PageVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageVo struct{}"
	}

	return strings.Join([]string{"PageVo", string(data)}, " ")
}
