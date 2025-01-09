package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandVolumeReq 扩容磁盘请求
type ExpandVolumeReq struct {

	// 订单ID，包周期桌面扩容时使用。
	OrderId *string `json:"order_id,omitempty"`

	// 扩容后的磁盘大小，单位为GB。
	NewSize int32 `json:"new_size"`
}

func (o ExpandVolumeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandVolumeReq struct{}"
	}

	return strings.Join([]string{"ExpandVolumeReq", string(data)}, " ")
}
