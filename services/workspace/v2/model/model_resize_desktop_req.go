package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopReq 变更规格请求。
type ResizeDesktopReq struct {

	// 桌面数据。支持批量将桌面变更为同一规格。
	Desktops []ResizeDesktopData `json:"desktops"`

	// 套餐id。批量变更时，则变更为同一规格的虚拟机。
	ProductId string `json:"product_id"`

	// 是否支持开机状态下执行变更规格操作。固定传值STOP_DESKTOP，如果桌面处于开机状态，会先关机再变更规格。
	Mode string `json:"mode"`

	// 是否自动放置，专属主机桌面变更规格时使用，默认是off关闭自动放置，on表示开启自动放置。
	AutoPlacement *string `json:"auto_placement,omitempty"`

	// 桌面池id。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`
}

func (o ResizeDesktopReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopReq struct{}"
	}

	return strings.Join([]string{"ResizeDesktopReq", string(data)}, " ")
}
