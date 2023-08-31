package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostClusterInfoDetail 主机集群信息
type HostClusterInfoDetail struct {

	// 主机集群id
	Id *string `json:"id,omitempty"`

	// 主机集群名
	Name *string `json:"name,omitempty"`

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

	// 是否是代理模式
	IsProxyMode *int32 `json:"is_proxy_mode,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o HostClusterInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostClusterInfoDetail struct{}"
	}

	return strings.Join([]string{"HostClusterInfoDetail", string(data)}, " ")
}
