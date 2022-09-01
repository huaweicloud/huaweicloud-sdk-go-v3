package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云手机使用的带宽信息。独占带宽按流量计费，共享带宽可选择带宽大小
type CreateCloudPhoneServerRequestBodyBandWidth struct {

	// 共享带宽ID，优先用该参数为云手机绑定带宽
	BandWidthId *string `json:"band_width_id,omitempty" xml:"band_width_id"`

	// band_width_id不存在时必选 - 0 表示独享带宽 - 1 表示共享带宽
	BandWidthShareType *int32 `json:"band_width_share_type,omitempty" xml:"band_width_share_type"`

	// 当band_width_share_type为共享带宽时必选 共享带宽默认取值范围5Mbit/s~2000Mbit/s 独享带宽的默认带宽是300Mbit/s
	BandWidthSize *int32 `json:"band_width_size,omitempty" xml:"band_width_size"`
}

func (o CreateCloudPhoneServerRequestBodyBandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneServerRequestBodyBandWidth struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneServerRequestBodyBandWidth", string(data)}, " ")
}
