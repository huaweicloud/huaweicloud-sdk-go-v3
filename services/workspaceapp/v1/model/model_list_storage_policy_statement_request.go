package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStoragePolicyStatementRequest Request Object
type ListStoragePolicyStatementRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListStoragePolicyStatementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStoragePolicyStatementRequest struct{}"
	}

	return strings.Join([]string{"ListStoragePolicyStatementRequest", string(data)}, " ")
}
