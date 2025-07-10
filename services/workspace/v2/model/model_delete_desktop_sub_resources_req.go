package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopSubResourcesReq 删除桌面协同资源请求体。
type DeleteDesktopSubResourcesReq struct {

	// 桌面ID列表。
	DesktopIds []string `json:"desktop_ids"`

	// 待删除附属资源类型。DESKTOP_SHARER：桌面协同资源。
	SubResourceType string `json:"sub_resource_type"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`
}

func (o DeleteDesktopSubResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopSubResourcesReq struct{}"
	}

	return strings.Join([]string{"DeleteDesktopSubResourcesReq", string(data)}, " ")
}
