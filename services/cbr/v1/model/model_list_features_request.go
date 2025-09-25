package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesRequest Request Object
type ListFeaturesRequest struct {

	// 每页显示条目数，正整数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移值,正整数
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListFeaturesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesRequest struct{}"
	}

	return strings.Join([]string{"ListFeaturesRequest", string(data)}, " ")
}
