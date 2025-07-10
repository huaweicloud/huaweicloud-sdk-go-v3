package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopTagsInfo 桌面携带标签数据结构。
type DesktopTagsInfo struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	// 携带标签列表。
	Tags []Tag `json:"tags"`
}

func (o DesktopTagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopTagsInfo struct{}"
	}

	return strings.Join([]string{"DesktopTagsInfo", string(data)}, " ")
}
