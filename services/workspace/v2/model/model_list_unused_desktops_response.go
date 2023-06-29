package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUnusedDesktopsResponse Response Object
type ListUnusedDesktopsResponse struct {

	// 指定时间段内未使用桌面列表。
	UnusedDesktops *[]UnusedDesktopInfo `json:"unused_desktops,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUnusedDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUnusedDesktopsResponse struct{}"
	}

	return strings.Join([]string{"ListUnusedDesktopsResponse", string(data)}, " ")
}
