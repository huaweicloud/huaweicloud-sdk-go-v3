package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopSubResourcesReq 桌面购买协同资源请求体。
type AddDesktopSubResourcesReq struct {

	// 桌面协同资源SKU码。
	SubResourceSku string `json:"sub_resource_sku"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 桌面ID列表。
	DesktopIds []string `json:"desktop_ids"`
}

func (o AddDesktopSubResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopSubResourcesReq struct{}"
	}

	return strings.Join([]string{"AddDesktopSubResourcesReq", string(data)}, " ")
}
