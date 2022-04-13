package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTagsByResourceTypeResponse struct {
	// 标签列表

	Tags *[]Tag `json:"tags,omitempty"`
	// 标签数量

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsByResourceTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsByResourceTypeResponse struct{}"
	}

	return strings.Join([]string{"ListTagsByResourceTypeResponse", string(data)}, " ")
}
