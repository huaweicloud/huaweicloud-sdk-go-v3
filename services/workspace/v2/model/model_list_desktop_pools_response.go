package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopPoolsResponse Response Object
type ListDesktopPoolsResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 桌面池信息。
	DesktopPools   *[]SimpleDesktopPoolInfo `json:"desktop_pools,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListDesktopPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopPoolsResponse", string(data)}, " ")
}
