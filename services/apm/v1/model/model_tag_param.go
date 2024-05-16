package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagParam struct {

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 关键字。
	Keyword *string `json:"keyword,omitempty"`

	// 是否分页。
	PageEnable *bool `json:"page_enable,omitempty"`

	// 每页容量。
	PageNumber *int32 `json:"page_number,omitempty"`

	// 当前页码。
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o TagParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagParam struct{}"
	}

	return strings.Join([]string{"TagParam", string(data)}, " ")
}
