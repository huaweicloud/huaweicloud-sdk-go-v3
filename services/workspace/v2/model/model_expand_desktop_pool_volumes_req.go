package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolVolumesReq 扩容桌面池磁盘请求。
type ExpandDesktopPoolVolumesReq struct {

	// 包周期订购ID，CBC订购回调时使用。
	OrderId *string `json:"order_id,omitempty"`

	// 扩容的桌面池磁盘列表。
	Volumes *[]VolumeInfo `json:"volumes,omitempty"`
}

func (o ExpandDesktopPoolVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolVolumesReq struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolVolumesReq", string(data)}, " ")
}
