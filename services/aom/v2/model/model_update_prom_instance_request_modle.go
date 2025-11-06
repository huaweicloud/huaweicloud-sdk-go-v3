package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePromInstanceRequestModle struct {

	// 待修改的普罗实例id。
	PromId string `json:"prom_id"`

	PromLimits *PromLimits `json:"prom_limits,omitempty"`

	// 待修改的普罗实例名称，名称不能以下划线或中划线开头和结尾，只含有中文，英文，数字，下划线，中划线，长度1-100。
	PromName *string `json:"prom_name,omitempty"`

	// 被聚合的账号和普罗实例列表。
	AggrPrometheusInfo *[]AggrPrometheusInfo `json:"aggr_prometheus_info,omitempty"`
}

func (o UpdatePromInstanceRequestModle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePromInstanceRequestModle struct{}"
	}

	return strings.Join([]string{"UpdatePromInstanceRequestModle", string(data)}, " ")
}
