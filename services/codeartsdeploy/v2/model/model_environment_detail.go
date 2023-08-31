package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvironmentDetail 环境详情
type EnvironmentDetail struct {

	// 环境id
	Id *string `json:"id,omitempty"`

	// 环境名称
	Name *string `json:"name,omitempty"`

	// 环境描述
	Description *string `json:"description,omitempty"`

	// 操作系统
	Os *string `json:"os,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 部署类型：0表示主机, 1表示kubernetes
	DeployType *int32 `json:"deploy_type,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 环境下主机实例数量
	InstanceCount *int32 `json:"instance_count,omitempty"`

	CreatedBy *UserInfo `json:"created_by,omitempty"`

	Permission *EnvironmentPermissionDetail `json:"permission,omitempty"`
}

func (o EnvironmentDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentDetail struct{}"
	}

	return strings.Join([]string{"EnvironmentDetail", string(data)}, " ")
}
