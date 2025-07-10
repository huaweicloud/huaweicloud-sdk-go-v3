package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopSidReq 更新桌面sid请求。
type UpdateDesktopSidReq struct {

	// 虚拟机列表。
	DesktopIds []string `json:"desktop_ids"`
}

func (o UpdateDesktopSidReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopSidReq struct{}"
	}

	return strings.Join([]string{"UpdateDesktopSidReq", string(data)}, " ")
}
