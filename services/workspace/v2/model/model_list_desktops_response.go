package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDesktopsResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 桌面信息。
	Desktops       *[]SimpleDesktopInfo `json:"desktops,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopsResponse", string(data)}, " ")
}
