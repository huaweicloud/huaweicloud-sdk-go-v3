package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckDesktopImagesReq struct {

	// 操作的桌面ID列表。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`
}

func (o CheckDesktopImagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDesktopImagesReq struct{}"
	}

	return strings.Join([]string{"CheckDesktopImagesReq", string(data)}, " ")
}
