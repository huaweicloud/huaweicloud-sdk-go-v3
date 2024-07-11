package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvironmentBaseInfo 环境基本信息
type EnvironmentBaseInfo struct {

	// 环境名称
	Name *string `json:"name,omitempty"`

	// 操作系统：windows|linux
	Os *string `json:"os,omitempty"`

	// 环境id
	Uuid *string `json:"uuid,omitempty"`

	// 主机集群id
	GroupId *string `json:"group_id,omitempty"`

	// 环境下主机数量
	HostCount *int32 `json:"host_count,omitempty"`
}

func (o EnvironmentBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentBaseInfo struct{}"
	}

	return strings.Join([]string{"EnvironmentBaseInfo", string(data)}, " ")
}
