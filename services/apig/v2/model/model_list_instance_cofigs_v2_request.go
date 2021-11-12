package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstanceCofigsV2Request struct {
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceCofigsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceCofigsV2Request struct{}"
	}

	return strings.Join([]string{"ListInstanceCofigsV2Request", string(data)}, " ")
}
