package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceCompliantRequest Request Object
type ListCceCompliantRequest struct {

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页限制
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCceCompliantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceCompliantRequest struct{}"
	}

	return strings.Join([]string{"ListCceCompliantRequest", string(data)}, " ")
}
