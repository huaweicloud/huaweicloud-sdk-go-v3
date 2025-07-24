package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareTypesRequest Request Object
type ListShareTypesRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 查询文件系统类型和配额返回的个数最大值
	Limit *int32 `json:"limit,omitempty"`

	// 查询文件系统类型和配额返回的偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListShareTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareTypesRequest struct{}"
	}

	return strings.Join([]string{"ListShareTypesRequest", string(data)}, " ")
}
