package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopDetachInfoResponse Response Object
type ListDesktopDetachInfoResponse struct {

	// 桌面解绑信息。
	DesktopDetachInfos *[]DesktopDetachInfo `json:"desktop_detach_infos,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListDesktopDetachInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopDetachInfoResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopDetachInfoResponse", string(data)}, " ")
}
