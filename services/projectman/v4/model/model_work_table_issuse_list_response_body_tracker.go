package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkTableIssuseListResponseBodyTracker 工作项类型信息
type WorkTableIssuseListResponseBodyTracker struct {

	// 工作项类型id
	Id *int32 `json:"id,omitempty"`

	// 工作项类型名称
	Name *string `json:"name,omitempty"`
}

func (o WorkTableIssuseListResponseBodyTracker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyTracker struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyTracker", string(data)}, " ")
}
