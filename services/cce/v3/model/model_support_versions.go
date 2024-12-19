package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportVersions 插件支持升级的集群版本
type SupportVersions struct {

	// 支持的集群类型
	ClusterType string `json:"clusterType"`

	// 支持的集群版本（正则表达式）
	ClusterVersion []string `json:"clusterVersion"`

	// 作用的集群类型 **取值范围：** - CCE：CCE Standard集群 - Turbo：CCE Turbo集群 - Autopilot：CCE Autopilot集群  **默认取值** 为空时默认为CCE Standard，CCE Turbo集群
	Category *[]string `json:"category,omitempty"`
}

func (o SupportVersions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportVersions struct{}"
	}

	return strings.Join([]string{"SupportVersions", string(data)}, " ")
}
