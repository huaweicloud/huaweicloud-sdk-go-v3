package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeVolume 包周期云硬盘退订的结果
type UnsubscribeVolume struct {

	// 卷id对应的结果
	VolumeId string `json:"volume_id"`

	// 卷id对应的退订订单id，如果是已到期的云硬盘退订，则不显示此字段。
	OrderId *string `json:"order_id,omitempty"`

	// volume_id对应的退订结果，只有SUCCESS 和 FAIL两种结果。
	Result string `json:"result"`

	// 当result为FAIL时，此字段显示具体的失败原因。 result为SUCCESS时，不显示此字段。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o UnsubscribeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeVolume struct{}"
	}

	return strings.Join([]string{"UnsubscribeVolume", string(data)}, " ")
}
