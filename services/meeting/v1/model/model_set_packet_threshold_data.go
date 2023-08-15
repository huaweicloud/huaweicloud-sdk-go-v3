package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPacketThresholdData 丢包率阈值设置请求体。
type SetPacketThresholdData struct {

	// 接收方向阈值设定值，单位为百分比(%)。 取值范围：0 - 100。
	Receiving *int32 `json:"receiving,omitempty"`

	// 发送方向阈值设定值，单位为百分比(%)。 取值范围：0 - 100。
	Sending *int32 `json:"sending,omitempty"`
}

func (o SetPacketThresholdData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPacketThresholdData struct{}"
	}

	return strings.Join([]string{"SetPacketThresholdData", string(data)}, " ")
}
