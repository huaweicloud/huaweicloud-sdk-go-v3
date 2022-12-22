package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSharesRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 返回的文件系统个数，最大值为200。
	Limit *int64 `json:"limit,omitempty"`

	// 文件系统查询个数的偏移量。
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListSharesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesRequest struct{}"
	}

	return strings.Join([]string{"ListSharesRequest", string(data)}, " ")
}
