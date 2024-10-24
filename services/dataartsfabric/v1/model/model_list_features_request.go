package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesRequest Request Object
type ListFeaturesRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFeaturesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesRequest struct{}"
	}

	return strings.Join([]string{"ListFeaturesRequest", string(data)}, " ")
}
