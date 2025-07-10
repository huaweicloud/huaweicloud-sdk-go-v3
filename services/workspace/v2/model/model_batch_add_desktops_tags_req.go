package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddDesktopsTagsReq 批量桌面添加标签。
type BatchAddDesktopsTagsReq struct {

	// 桌面（桌面携带标签列表）列表。批量操作时非法桌面ID会过滤不做处理。
	Desktops []DesktopTagsInfo `json:"desktops"`
}

func (o BatchAddDesktopsTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDesktopsTagsReq struct{}"
	}

	return strings.Join([]string{"BatchAddDesktopsTagsReq", string(data)}, " ")
}
