package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSpecsRequest struct {

	// 规格编码
	SpecCode *string `json:"spec_code,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询个数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecsRequest struct{}"
	}

	return strings.Join([]string{"ListSpecsRequest", string(data)}, " ")
}
