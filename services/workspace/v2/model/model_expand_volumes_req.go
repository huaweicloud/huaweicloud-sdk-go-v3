package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandVolumesReq 单个桌面扩容磁盘参数。
type ExpandVolumesReq struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 订单ID，包周期桌面扩容时使用。
	OrderId *string `json:"order_id,omitempty"`

	// 磁盘ID。
	VolumeId *string `json:"volume_id,omitempty"`

	// 扩容后的磁盘大小，单位为GB。
	NewSize *int32 `json:"new_size,omitempty"`
}

func (o ExpandVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandVolumesReq struct{}"
	}

	return strings.Join([]string{"ExpandVolumesReq", string(data)}, " ")
}
