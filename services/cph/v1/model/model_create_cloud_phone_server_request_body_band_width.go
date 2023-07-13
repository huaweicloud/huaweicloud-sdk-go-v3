package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudPhoneServerRequestBodyBandWidth 云手机使用的带宽信息。独占带宽按流量计费，共享带宽可选择带宽大小
type CreateCloudPhoneServerRequestBodyBandWidth struct {

	// 云手机服务器的带宽唯一标识
	BandWidthId *string `json:"band_width_id,omitempty"`

	// 云手机服务器的带宽大小
	BandWidthSize *int32 `json:"band_width_size,omitempty"`

	// 云手机服务器的带宽类型  - 0，per，独享带宽 - 1，whole，共享带宽
	BandWidthShareType *int32 `json:"band_width_share_type,omitempty"`
}

func (o CreateCloudPhoneServerRequestBodyBandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneServerRequestBodyBandWidth struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneServerRequestBodyBandWidth", string(data)}, " ")
}
