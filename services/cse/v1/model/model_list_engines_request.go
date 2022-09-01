package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEnginesRequest struct {

	// 偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量。
	Limit *string `json:"limit,omitempty" xml:"limit"`
}

func (o ListEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesRequest struct{}"
	}

	return strings.Join([]string{"ListEnginesRequest", string(data)}, " ")
}
