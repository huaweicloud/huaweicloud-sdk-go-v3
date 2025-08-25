package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudServiceInfo struct {

	// 当前租户开通的云原生密码服务数量（只统计白名单服务情况）
	ServiceNum int32 `json:"service_num"`

	// 当前租户云原生密码服务的资源实例总和
	ResourceNum int32 `json:"resource_num"`

	ResourceDistribution *ResourceDistribution `json:"resource_distribution"`
}

func (o CloudServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudServiceInfo struct{}"
	}

	return strings.Join([]string{"CloudServiceInfo", string(data)}, " ")
}
