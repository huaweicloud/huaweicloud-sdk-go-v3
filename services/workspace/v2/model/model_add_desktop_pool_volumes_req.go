package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopPoolVolumesReq 扩容桌面池请求。
type AddDesktopPoolVolumesReq struct {

	// 包周期订购ID，CBC订购回调时使用。
	OrderId *string `json:"order_id,omitempty"`

	// 增加的磁盘列表。
	Volumes *[]VolumeAddInfo `json:"volumes,omitempty"`
}

func (o AddDesktopPoolVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopPoolVolumesReq struct{}"
	}

	return strings.Join([]string{"AddDesktopPoolVolumesReq", string(data)}, " ")
}
