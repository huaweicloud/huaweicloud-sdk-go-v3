package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListObsBucketsRequest struct {

	// 分页查询时的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页一页显示数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListObsBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListObsBucketsRequest", string(data)}, " ")
}
