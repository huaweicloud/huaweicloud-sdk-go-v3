package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopVersionResponse Response Object
type ListDesktopVersionResponse struct {

	// 应用对象列表。
	Desktops *[]DesktopVersionInfo `json:"desktops,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopVersionResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopVersionResponse", string(data)}, " ")
}
