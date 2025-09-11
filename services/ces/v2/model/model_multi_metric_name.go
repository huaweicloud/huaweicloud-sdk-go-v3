package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiMetricName **参数解释** 多个指标名称 **约束限制** 不涉及 **取值范围** 长度为[1,1080]个字符，多个指标名称之间用逗号隔开 **默认取值** 不涉及
type MultiMetricName struct {
}

func (o MultiMetricName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiMetricName struct{}"
	}

	return strings.Join([]string{"MultiMetricName", string(data)}, " ")
}
