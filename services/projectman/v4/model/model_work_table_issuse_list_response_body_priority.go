package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkTableIssuseListResponseBodyPriority 工作项优先级信息
type WorkTableIssuseListResponseBodyPriority struct {

	// 工作项优先级id
	Id *int32 `json:"id,omitempty"`

	// 工作项优先级名称
	Name *string `json:"name,omitempty"`
}

func (o WorkTableIssuseListResponseBodyPriority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyPriority struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyPriority", string(data)}, " ")
}
