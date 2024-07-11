package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostClusterInfo 主机集群信息
type HostClusterInfo struct {

	// 主机集群id
	Id *string `json:"id,omitempty"`

	// 集群内主机数量，一个主机集群内最多可添加200台主机
	HostCount *int32 `json:"host_count,omitempty"`

	// 主机集群名
	Name *string `json:"name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 操作系统：windows|linux
	Os *string `json:"os,omitempty"`

	// slave集群id，默认为null时使用默认slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	CreatedBy *UserInfo `json:"created_by,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	Permission *PermissionClusterDetail `json:"permission,omitempty"`

	// 创建人名称
	NickName *string `json:"nick_name,omitempty"`

	// 环境数量
	EnvCount *int32 `json:"env_count,omitempty"`
}

func (o HostClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostClusterInfo struct{}"
	}

	return strings.Join([]string{"HostClusterInfo", string(data)}, " ")
}
