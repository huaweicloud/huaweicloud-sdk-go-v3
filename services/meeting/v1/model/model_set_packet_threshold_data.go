package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 丢包率阈值设置请求体
type SetPacketThresholdData struct {

	// 接收方向阈值设定值，单位为百分比(%)。 取值范围：0 - 100。
	Receiving *int32 `json:"receiving,omitempty" xml:"receiving"`

	// 发送方向阈值设定值，单位为百分比(%)。 取值范围：0 - 100。
	Sending *int32 `json:"sending,omitempty" xml:"sending"`
}

func (o SetPacketThresholdData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPacketThresholdData struct{}"
	}

	return strings.Join([]string{"SetPacketThresholdData", string(data)}, " ")
}
