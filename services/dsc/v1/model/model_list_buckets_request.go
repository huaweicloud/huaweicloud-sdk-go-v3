package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBucketsRequest Request Object
type ListBucketsRequest struct {

	// 已授权
	Added *bool `json:"added,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListBucketsRequest", string(data)}, " ")
}
