package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopResourceItem 桌面资源包。
type DesktopResourceItem struct {

	// 资源包名称。
	ResourcePackageName string `json:"resource_package_name"`

	// 桌面产品编码。
	DesktopResourceSpecCode string `json:"desktop_resource_spec_code"`

	// 资源包产品编码。
	ResourceSpecCode string `json:"resource_spec_code"`

	// 生效时间，格式：yyyy-MM-ddTHH:mm:ssZ（2025-04-12T17:30:00Z）。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 是否自动续订。
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`
}

func (o DesktopResourceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopResourceItem struct{}"
	}

	return strings.Join([]string{"DesktopResourceItem", string(data)}, " ")
}
