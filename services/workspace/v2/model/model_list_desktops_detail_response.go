package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDesktopsDetailResponse struct {

	// 桌面详情列表。
	Desktops *[]DesktopDetailInfo `json:"desktops,omitempty"`

	// 桌面总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopsDetailResponse", string(data)}, " ")
}
