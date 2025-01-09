package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopActionsResponse Response Object
type ListDesktopActionsResponse struct {

	// 错误码，失败时返回。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 桌面开关列表
	DesktopActions *[]DesktopAction `json:"desktop_actions,omitempty"`

	// 总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopActionsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopActionsResponse", string(data)}, " ")
}
