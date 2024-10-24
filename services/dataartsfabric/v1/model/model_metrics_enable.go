package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricsEnable 是否开启AOM监控采集。
type MetricsEnable struct {
}

func (o MetricsEnable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsEnable struct{}"
	}

	return strings.Join([]string{"MetricsEnable", string(data)}, " ")
}
