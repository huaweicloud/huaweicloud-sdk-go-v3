package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerRequestBodyBandWidth 云手机使用的带宽信息。
type ChangeCloudPhoneServerRequestBodyBandWidth struct {

	// 云手机服务器的带宽唯一标识。
	BandWidthId *string `json:"band_width_id,omitempty"`

	// 云手机服务器的带宽大小。
	BandWidthSize *int32 `json:"band_width_size,omitempty"`

	// 云手机服务器带宽的计费方式。  - 0：bandwidth, 按带宽计费  - 1：traffic, 按流量计费
	BandWidthChargeMode int32 `json:"band_width_charge_mode"`

	// 云手机服务器的带宽类型。  - 0：per，独享带宽 - 1：whole，共享带宽
	BandWidthShareType int32 `json:"band_width_share_type"`
}

func (o ChangeCloudPhoneServerRequestBodyBandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerRequestBodyBandWidth struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerRequestBodyBandWidth", string(data)}, " ")
}
