package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppDetailInfo 应用详情信息
type AppDetailInfo struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用所属区域
	Region *string `json:"region,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 当前应用是否被禁用
	IsDisable *bool `json:"is_disable,omitempty"`

	// 创建方式
	CreateType *string `json:"create_type,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// slave集群id，默认为null时使用默认slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	// 当前用户是否已收藏
	IsCare *bool `json:"is_care,omitempty"`

	// 是否有编辑权限
	CanModify *bool `json:"can_modify,omitempty"`

	// 是否有删除的权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有部署权限
	CanExecute *bool `json:"can_execute,omitempty"`

	// 是否有复制权限
	CanCopy *bool `json:"can_copy,omitempty"`

	// 是否有编辑应用权限矩阵的权限
	CanManage *bool `json:"can_manage,omitempty"`

	// 是否有创建环境的权限
	CanCreateEnv *bool `json:"can_create_env,omitempty"`

	// 是否有禁用应用的权限
	CanDisable *bool `json:"can_disable,omitempty"`

	// 应用所属人租户id
	OwnerTenantId *string `json:"owner_tenant_id,omitempty"`

	// 应用创建者用户名
	CreateUserId *string `json:"create_user_id,omitempty"`

	// 应用创建人租户id
	CreateTenantId *string `json:"create_tenant_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 权限等级
	PermissionLevel *string `json:"permission_level,omitempty"`

	// 部署任务信息
	ArrangeInfos *[]TaskV2Info `json:"arrange_infos,omitempty"`
}

func (o AppDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppDetailInfo struct{}"
	}

	return strings.Join([]string{"AppDetailInfo", string(data)}, " ")
}
