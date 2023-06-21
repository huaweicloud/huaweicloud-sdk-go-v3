package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取租户的习题库调用参数
type PackagesListRequestBody struct {
	Filter *PackageFilter `json:"filter,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 起始页
	StartIndex *int32 `json:"start_index,omitempty"`
}

func (o PackagesListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackagesListRequestBody struct{}"
	}

	return strings.Join([]string{"PackagesListRequestBody", string(data)}, " ")
}
