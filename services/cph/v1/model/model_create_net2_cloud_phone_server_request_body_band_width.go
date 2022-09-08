package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云手机使用的带宽信息
type CreateNet2CloudPhoneServerRequestBodyBandWidth struct {

	// 带宽类型 - 0 表示独享带宽 - 1 表示共享带宽
	BandWidthShareType int32 `json:"band_width_share_type"`

	// 功能说明：带宽大小  带宽（Mbit/s），取值范围为[1,2000]。  调整带宽时的最小单位会根据带宽范围不同存在差异。  小于等于300Mbit/s：默认最小单位为1Mbit/s。 300Mbit/s~1000Mbit/s：默认最小单位为50Mbit/s。 大于1000Mbit/s：默认最小单位为500Mbit/s。 说明：  如果share_type是独享带宽，该参数必选项；如果share_type是共享带宽并且id有值，该参数会忽略
	BandWidthSize *int32 `json:"band_width_size,omitempty"`

	// 带宽ID，创建共享带宽类型带宽的弹性IP时可以指定之前的共享带宽创建。  取值范围：共享带宽类型的带宽ID。  说明：  当创建共享带宽类型的带宽时，该字段必选
	BandWidthId *string `json:"band_width_id,omitempty"`
}

func (o CreateNet2CloudPhoneServerRequestBodyBandWidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyBandWidth struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyBandWidth", string(data)}, " ")
}
