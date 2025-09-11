package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiMetricNameResp **参数解释** 多个指标名称 **取值范围** 长度为[1,1080]个字符，多个指标名称之间用逗号隔开
type MultiMetricNameResp struct {
}

func (o MultiMetricNameResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiMetricNameResp struct{}"
	}

	return strings.Join([]string{"MultiMetricNameResp", string(data)}, " ")
}
