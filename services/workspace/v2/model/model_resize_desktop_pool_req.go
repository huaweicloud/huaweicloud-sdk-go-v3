package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopPoolReq 变更桌面池请求。
type ResizeDesktopPoolReq struct {

	// 目标规格产品ID。
	ProductId string `json:"product_id"`

	// 产品规格ID。可用区是边缘可用区时，必填此参数。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 是否支持开机状态下执行变更规格操作。固定传值STOP_DESKTOP，如果桌面处于开机状态，会先关机再变更规格。
	Mode *string `json:"mode,omitempty"`
}

func (o ResizeDesktopPoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopPoolReq struct{}"
	}

	return strings.Join([]string{"ResizeDesktopPoolReq", string(data)}, " ")
}
