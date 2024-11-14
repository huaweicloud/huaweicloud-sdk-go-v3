package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharesRequest Request Object
type ListSharesRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 设置返回的文件系统个数的最大值，不填默认为1000个
	Limit *int64 `json:"limit,omitempty"`

	// 设置返回的文件系统的偏移量。
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListSharesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesRequest struct{}"
	}

	return strings.Join([]string{"ListSharesRequest", string(data)}, " ")
}
