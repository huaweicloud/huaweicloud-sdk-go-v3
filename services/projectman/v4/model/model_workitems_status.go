package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项状态
type WorkitemsStatus struct {

	// 状态id
	Id *string `json:"id,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`
}

func (o WorkitemsStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemsStatus struct{}"
	}

	return strings.Join([]string{"WorkitemsStatus", string(data)}, " ")
}
