package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsedDesktopInfoResponse Response Object
type ListUsedDesktopInfoResponse struct {

	// 桌面使用信息（以桌面Id划分）。
	DesktopUsedInfoList *[]DesktopUsedHoursInfo `json:"desktop_used_info_list,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUsedDesktopInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsedDesktopInfoResponse struct{}"
	}

	return strings.Join([]string{"ListUsedDesktopInfoResponse", string(data)}, " ")
}
