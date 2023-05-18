package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项状态信息
type WorkTableIssuseListResponseBodyStatus struct {

	// 工作项状态id
	Id *int32 `json:"id,omitempty"`

	// 工作项优先级名称
	Name *string `json:"name,omitempty"`
}

func (o WorkTableIssuseListResponseBodyStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyStatus struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyStatus", string(data)}, " ")
}
