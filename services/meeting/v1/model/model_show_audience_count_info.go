package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAudienceCountInfo struct {

	// 观众显示策略：服务端用来计算观众人数、下发给客户端控制观众显示 * 0：不显示 * 1：倍增显示与会人数，基于实时与会人数或累计与会人次（假设为N），可以再进行倍增设置。 * ---支持设置倍增倍数X和基础人数Y，设置后，显示的人数为：N*X+Y。 * ---X支持设置到小数点后1位，当N*X计算非整数时，向下取整。X的范围是1~10，Y的范围是0~10000。 * 2：按累计与会人数显示（当前不做）
	ShowAudienceMode *int32 `json:"showAudienceMode,omitempty"`

	// 基础人数，范围是0~10000
	BaseAudienceCount *int32 `json:"baseAudienceCount,omitempty"`

	// 倍增倍数，范围是1~10, 支持设置到小数点后1位
	Multiple *float64 `json:"multiple,omitempty"`
}

func (o ShowAudienceCountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAudienceCountInfo struct{}"
	}

	return strings.Join([]string{"ShowAudienceCountInfo", string(data)}, " ")
}
