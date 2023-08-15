package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PointValidityingDto 点位校验配置结构体
type PointValidityingDto struct {

	// 点位上报值的最小值，小于该值则上报告警
	Min int64 `json:"min"`

	// 点位上报值的最大值，大于该值则上报告警
	Max int64 `json:"max"`
}

func (o PointValidityingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PointValidityingDto struct{}"
	}

	return strings.Join([]string{"PointValidityingDto", string(data)}, " ")
}
