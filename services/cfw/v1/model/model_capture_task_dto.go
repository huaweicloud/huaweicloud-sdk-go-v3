package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaptureTaskDto struct {
	Destination *CaptureRuleAddressDto `json:"destination"`

	// 抓包时长
	Duration int32 `json:"duration"`

	// 最大抓包数
	MaxPackets int32 `json:"max_packets"`

	// 抓包任务名称
	Name string `json:"name"`

	Service *CaptureServiceDto `json:"service"`

	Source *CaptureRuleAddressDto `json:"source"`
}

func (o CaptureTaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaptureTaskDto struct{}"
	}

	return strings.Join([]string{"CaptureTaskDto", string(data)}, " ")
}
