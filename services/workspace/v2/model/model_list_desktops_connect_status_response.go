package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsConnectStatusResponse Response Object
type ListDesktopsConnectStatusResponse struct {

	// 桌面登录信息列表。
	Desktops *[]ConnectDesktopsInfo `json:"desktops,omitempty"`

	// 桌面总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopsConnectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsConnectStatusResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopsConnectStatusResponse", string(data)}, " ")
}
