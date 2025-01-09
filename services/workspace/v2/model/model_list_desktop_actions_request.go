package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopActionsRequest Request Object
type ListDesktopActionsRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	// 每页限制数
	Offset *int32 `json:"offset,omitempty"`

	// 起始位置
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopActionsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopActionsRequest", string(data)}, " ")
}
