package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildDesktopPoolReq 重建系统盘的请求。
type RebuildDesktopPoolReq struct {

	// 镜像类型。
	ImageType *string `json:"image_type,omitempty"`

	// 模板ID。
	ImageId string `json:"image_id"`

	// os类型。
	OsType *string `json:"os_type,omitempty"`

	// 立即重建时给用户预留的保存数据的时间（单位：分钟）。
	DelayTime *int32 `json:"delay_time,omitempty"`

	// 下发重建系统盘任务时，给用户发送的提示信息。
	Message *string `json:"message,omitempty"`

	// 订单ID，包周期桌面重建系统盘，涉及收费镜像时需传
	OrderId *string `json:"order_id,omitempty"`

	// 是否是修复行为，修复行为只修复镜像ID与桌面池镜像ID不一致的桌面，用于桌面池切换镜像失败场景的修复。
	IsFix *bool `json:"is_fix,omitempty"`
}

func (o RebuildDesktopPoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildDesktopPoolReq struct{}"
	}

	return strings.Join([]string{"RebuildDesktopPoolReq", string(data)}, " ")
}
