package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildDesktopsReq 重建系统盘的请求。
type RebuildDesktopsReq struct {

	// 计算机id列表。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

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

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o RebuildDesktopsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildDesktopsReq struct{}"
	}

	return strings.Join([]string{"RebuildDesktopsReq", string(data)}, " ")
}
