package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchoverRatioInfo struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 容灾切换的故障节点比例，下限是50，步长是10，最大是100，默认为100。
	SwitchoverRatio *int32 `json:"switchover_ratio,omitempty"`
}

func (o SwitchoverRatioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverRatioInfo struct{}"
	}

	return strings.Join([]string{"SwitchoverRatioInfo", string(data)}, " ")
}
