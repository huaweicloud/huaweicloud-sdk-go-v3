package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreUpgradeProbeDto struct {

	// 端口
	Port *int32 `json:"port,omitempty"`

	// 请求路径
	Path *string `json:"path,omitempty"`

	// 轮询间隔
	Interval *int32 `json:"interval,omitempty"`

	// 协议类型
	Protocol *string `json:"protocol,omitempty"`

	TimeoutConfig *UpgradeProbeTimeoutConfigDto `json:"timeout_config,omitempty"`
}

func (o PreUpgradeProbeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreUpgradeProbeDto struct{}"
	}

	return strings.Join([]string{"PreUpgradeProbeDto", string(data)}, " ")
}
