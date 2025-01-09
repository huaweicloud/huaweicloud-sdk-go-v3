package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopReq 变更规格请求。
type ResizeDesktopReq struct {

	// 桌面数据。支持批量按需类型桌面变更为同一规格。
	Desktops []ResizeDesktopData `json:"desktops"`

	// 套餐id。批量变更时，则变更为同一规格的虚拟机。
	ProductId string `json:"product_id"`

	// 套餐flavorId。批量变更时，则变更为同一规格的虚拟机。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 是否支持开机状态下执行变更规格操作。固定传值STOP_DESKTOP，如果桌面处于开机状态，会先关机再变更规格。
	Mode string `json:"mode"`

	// 专属主机ID。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	// 订单ID，包周期变更规格时使用。
	OrderId *string `json:"order_id,omitempty"`

	ExtendParam *ResizeDesktopExtendParam `json:"extend_param,omitempty"`
}

func (o ResizeDesktopReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopReq struct{}"
	}

	return strings.Join([]string{"ResizeDesktopReq", string(data)}, " ")
}
