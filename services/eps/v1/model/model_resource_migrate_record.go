package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceMigrateRecord struct {

	// 是否关联迁移
	Associated *bool `json:"associated,omitempty"`

	// 响应码
	Code *string `json:"code,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 事件时间
	EventTime *string `json:"event_time,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 迁移类型
	OperateType *string `json:"operate_type,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 发起迁移的企业项目ID
	InitiateEpId *string `json:"initiate_ep_id,omitempty"`

	// 发起迁移的业项目名称
	InitiateEpName *string `json:"initiate_ep_name,omitempty"`

	// 源企业项目ID
	OriginEpId *string `json:"origin_ep_id,omitempty"`

	// 源企业项目名称
	OriginEpName *string `json:"origin_ep_name,omitempty"`

	// 目标企业项目ID
	TargetEpId *string `json:"target_ep_id,omitempty"`

	// 目标企业项目名称
	TargetEpName *string `json:"target_ep_name,omitempty"`

	// 是否存在失败
	ExistFailed *string `json:"exist_failed,omitempty"`
}

func (o ResourceMigrateRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceMigrateRecord struct{}"
	}

	return strings.Join([]string{"ResourceMigrateRecord", string(data)}, " ")
}
