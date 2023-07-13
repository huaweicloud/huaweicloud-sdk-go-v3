package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCategoryRequest Request Object
type ListCategoryRequest struct {

	// workspace 信息
	Workspace string `json:"workspace"`

	// 分页时每页的条数,最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`

	// application/json
	Accept string `json:"accept"`
}

func (o ListCategoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCategoryRequest struct{}"
	}

	return strings.Join([]string{"ListCategoryRequest", string(data)}, " ")
}
