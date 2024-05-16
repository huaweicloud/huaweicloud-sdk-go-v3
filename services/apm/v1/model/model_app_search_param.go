package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppSearchParam 应用查询返回的数据结构。
type AppSearchParam struct {

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 区域名称。
	Region string `json:"region"`

	// 页码。
	Page int32 `json:"page"`

	// 每页条数。
	PageSize *int32 `json:"page_size,omitempty"`

	// 关键字。
	Keyword *string `json:"keyword,omitempty"`
}

func (o AppSearchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppSearchParam struct{}"
	}

	return strings.Join([]string{"AppSearchParam", string(data)}, " ")
}
