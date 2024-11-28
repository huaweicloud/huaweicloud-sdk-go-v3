package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopPoolVolumesReq 删除桌面池磁盘请求。
type DeleteDesktopPoolVolumesReq struct {

	// 删除的桌面池磁盘列表。
	Volumes *[]VolumeInfo `json:"volumes,omitempty"`
}

func (o DeleteDesktopPoolVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopPoolVolumesReq struct{}"
	}

	return strings.Join([]string{"DeleteDesktopPoolVolumesReq", string(data)}, " ")
}
