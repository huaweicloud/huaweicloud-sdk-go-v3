package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DimensionNames 启用一键告警关联告警规则的维度列表，包括指标告警&事件告警，且至少开启一个
type DimensionNames struct {

	// 启用一键告警关联指标告警规则的维度列表，未指定的维度默认不开启
	Metric []string `json:"metric"`

	// 启用一键告警关联事件告警规则的维度列表，未指定的维度默认不开启，其中\"\"代表全部资源
	Event []string `json:"event"`
}

func (o DimensionNames) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionNames struct{}"
	}

	return strings.Join([]string{"DimensionNames", string(data)}, " ")
}
