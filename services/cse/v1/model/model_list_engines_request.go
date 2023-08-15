package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnginesRequest Request Object
type ListEnginesRequest struct {

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *string `json:"limit,omitempty"`
}

func (o ListEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesRequest struct{}"
	}

	return strings.Join([]string{"ListEnginesRequest", string(data)}, " ")
}
