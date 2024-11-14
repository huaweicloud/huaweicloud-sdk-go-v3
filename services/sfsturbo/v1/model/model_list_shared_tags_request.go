package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharedTagsRequest Request Object
type ListSharedTagsRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 返回的标签个数
	Limit *int32 `json:"limit,omitempty"`

	// 标签查询个数的偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSharedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSharedTagsRequest", string(data)}, " ")
}
