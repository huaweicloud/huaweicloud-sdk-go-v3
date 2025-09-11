package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoVo struct {

	// 页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 单页大小
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o PageInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoVo struct{}"
	}

	return strings.Join([]string{"PageInfoVo", string(data)}, " ")
}
