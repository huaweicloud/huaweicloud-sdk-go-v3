package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiMetricName 多个指标名称，用逗号隔开
type MultiMetricName struct {
}

func (o MultiMetricName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiMetricName struct{}"
	}

	return strings.Join([]string{"MultiMetricName", string(data)}, " ")
}
