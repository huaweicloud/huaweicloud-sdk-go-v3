package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeProbeTimeoutConfigDto struct {

	// 超时时间
	Timeout *int32 `json:"timeout,omitempty"`

	// 失败阈值
	FailureThreshold *int32 `json:"failure_threshold,omitempty"`
}

func (o UpgradeProbeTimeoutConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeProbeTimeoutConfigDto struct{}"
	}

	return strings.Join([]string{"UpgradeProbeTimeoutConfigDto", string(data)}, " ")
}
