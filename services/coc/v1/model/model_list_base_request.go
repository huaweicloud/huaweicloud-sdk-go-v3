package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaseRequest limit   最小值0 最大值Integer.MAX offset 查询数量  值域 0 到1000
type ListBaseRequest struct {

	// limit
	Limit *int64 `json:"limit,omitempty"`

	// 查询数量 最小值0 最大值1000
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListBaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaseRequest struct{}"
	}

	return strings.Join([]string{"ListBaseRequest", string(data)}, " ")
}
