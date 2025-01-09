package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopVolumesReq 单个桌面添加磁盘参数。
type AddDesktopVolumesReq struct {

	// 桌面Id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 订单ID，包周期桌面添加磁盘时使用。
	OrderId *string `json:"order_id,omitempty"`

	// 待新增的磁盘信息，每个桌面的数据盘数量不超过10个。
	Volumes *[]Volume `json:"volumes,omitempty"`
}

func (o AddDesktopVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopVolumesReq struct{}"
	}

	return strings.Join([]string{"AddDesktopVolumesReq", string(data)}, " ")
}
