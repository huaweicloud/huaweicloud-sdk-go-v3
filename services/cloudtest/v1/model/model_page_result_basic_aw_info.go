package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageResultBasicAwInfo struct {

	// 当前页数据
	PageList *[]BasicAwInfo `json:"page_list,omitempty"`

	// 当前页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 总页数
	TotalPage *int32 `json:"total_page,omitempty"`

	// 总条数
	TotalSize *int64 `json:"total_size,omitempty"`
}

func (o PageResultBasicAwInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageResultBasicAwInfo struct{}"
	}

	return strings.Join([]string{"PageResultBasicAwInfo", string(data)}, " ")
}
