package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaseModelsRequest Request Object
type ListBaseModelsRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListBaseModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaseModelsRequest struct{}"
	}

	return strings.Join([]string{"ListBaseModelsRequest", string(data)}, " ")
}
