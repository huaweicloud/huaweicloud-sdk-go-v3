package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Bandwidth 带宽信息响应
type Bandwidth struct {

	// 云手机服务器的带宽名称
	BandWidthName *string `json:"band_width_name,omitempty"`

	// 云手机服务器的带宽唯一标识
	BandWidthId *string `json:"band_width_id,omitempty"`

	// 云手机服务器的带宽大小
	BandWidthSize *int32 `json:"band_width_size,omitempty"`

	// 云手机服务器带宽的计费方式  取值范围：  - 0，bandwidth, 按带宽计费  - 1，traffic, 按流量计费
	BandWidthChargeMode *int32 `json:"band_width_charge_mode,omitempty"`

	// 云手机服务器的带宽类型  - 0，per，独享带宽 - 1，whole，共享带宽
	BandWidthShareType *int32 `json:"band_width_share_type,omitempty"`

	// 带宽创建时间  时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	CreateTime *string `json:"create_time,omitempty"`

	// 带宽更新时间  时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Bandwidth struct{}"
	}

	return strings.Join([]string{"Bandwidth", string(data)}, " ")
}
