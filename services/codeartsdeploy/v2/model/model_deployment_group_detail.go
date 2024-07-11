package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeploymentGroupDetail 主机集群详情响应体
type DeploymentGroupDetail struct {

	// 主机集群id
	GroupId *string `json:"group_id,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 修改时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 集群内主机数量，一个主机集群内最多可添加200台主机
	HostCount *int32 `json:"host_count,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 主机集群名
	Name *string `json:"name,omitempty"`

	// 局点信息
	RegionName *string `json:"region_name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 操作系统：windows|linux
	Os *string `json:"os,omitempty"`

	// 自动测试功能已下架，该字段已失效
	AutoConnectionTestSwitch *int32 `json:"auto_connection_test_switch,omitempty"`

	// slave集群id，默认为null时使用默认slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	CreatedBy *UserInfo `json:"created_by,omitempty"`

	UpdatedBy *UserInfo `json:"updated_by,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	Permission *PermissionGroupDetail `json:"permission,omitempty"`
}

func (o DeploymentGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentGroupDetail struct{}"
	}

	return strings.Join([]string{"DeploymentGroupDetail", string(data)}, " ")
}
