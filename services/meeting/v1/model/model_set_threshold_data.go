package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 阈值阈值设置请求体
type SetThresholdData struct {

	// 接收方向阈值设定值，单位为毫秒(ms) 取值范围：0 - 10000
	Receiving *int32 `json:"receiving,omitempty"`

	// 发送方向阈值设定值，单位为毫秒(ms) 取值范围：0 - 10000
	Sending *int32 `json:"sending,omitempty"`
}

func (o SetThresholdData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetThresholdData struct{}"
	}

	return strings.Join([]string{"SetThresholdData", string(data)}, " ")
}
