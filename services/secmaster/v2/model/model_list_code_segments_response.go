package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCodeSegmentsResponse Response Object
type ListCodeSegmentsResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 代码片段
	Records        *[]CodeSegment `json:"records,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCodeSegmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCodeSegmentsResponse struct{}"
	}

	return strings.Join([]string{"ListCodeSegmentsResponse", string(data)}, " ")
}
