package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主机组详情响应体
type DeploymentGroupDetail struct {

	// 主机组id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 修改时间
	UpdatedTime *string `json:"updated_time,omitempty" xml:"updated_time"`

	// 组内主机数量，一个主机组内最多可添加200台主机
	HostCount *int32 `json:"host_count,omitempty" xml:"host_count"`

	// devcloud项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 主机组名
	Name *string `json:"name,omitempty" xml:"name"`

	// 局点信息
	RegionName *string `json:"region_name,omitempty" xml:"region_name"`

	// devcloud项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 操作系统：windows|linux
	Os *string `json:"os,omitempty" xml:"os"`

	// 自动连通性验证 0不执行 1每日 2每周
	AutoConnectionTestSwitch *int32 `json:"auto_connection_test_switch,omitempty" xml:"auto_connection_test_switch"`

	// slave集群id，默认为null时使用devcloud八爪鱼slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty" xml:"slave_cluster_id"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	CreatedBy *UserInfo `json:"created_by,omitempty" xml:"created_by"`

	UpdateBy *UserInfo `json:"update_by,omitempty" xml:"update_by"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	Permission *PermissionGroupDetail `json:"permission,omitempty" xml:"permission"`
}

func (o DeploymentGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentGroupDetail struct{}"
	}

	return strings.Join([]string{"DeploymentGroupDetail", string(data)}, " ")
}
