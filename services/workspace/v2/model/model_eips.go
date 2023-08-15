package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Eips 桌面EIP。
type Eips struct {

	// EIP的id。
	Id *string `json:"id,omitempty"`

	// Eip地址。
	Address *string `json:"address,omitempty"`

	// 带宽大小。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// traffic（按流量计费），bandwidth（按带宽计费）。
	EipChargeMode *string `json:"eip_charge_mode,omitempty"`

	// 创建时间，格式为：yyyy-MM-ddTHH:mm:ssZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 绑定的桌面id。
	AttachedDesktopId *string `json:"attached_desktop_id,omitempty"`

	// 绑定的桌面名称。
	AttachedDesktopName *string `json:"attached_desktop_name,omitempty"`
}

func (o Eips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Eips struct{}"
	}

	return strings.Join([]string{"Eips", string(data)}, " ")
}
