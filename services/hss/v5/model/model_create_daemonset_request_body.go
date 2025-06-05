package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDaemonsetRequestBody struct {

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 开启agent自动升级
	AutoUpgrade *bool `json:"auto_upgrade,omitempty"`

	// 容器运行时配置
	RuntimeInfo *[]RuntimeRequestBody `json:"runtime_info,omitempty"`

	ScheduleInfo *CreateDaemonsetRequestBodyScheduleInfo `json:"schedule_info,omitempty"`
}

func (o CreateDaemonsetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDaemonsetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDaemonsetRequestBody", string(data)}, " ")
}
