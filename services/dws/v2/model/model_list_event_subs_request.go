package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSubsRequest Request Object
type ListEventSubsRequest struct {

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 限制条目数
	Limit *string `json:"limit,omitempty"`
}

func (o ListEventSubsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSubsRequest struct{}"
	}

	return strings.Join([]string{"ListEventSubsRequest", string(data)}, " ")
}
