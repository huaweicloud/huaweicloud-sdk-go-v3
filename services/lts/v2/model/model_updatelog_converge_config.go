package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatelogConvergeConfig 组织成员日志汇聚配置
type UpdatelogConvergeConfig struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 组织成员账号id
	MemberAccountId string `json:"member_account_id"`

	// 组织成员项目id
	MemberProjectId *string `json:"member_project_id,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 组织id
	OrganizationId string `json:"organization_id"`

	// 管理员或者委托管理员账号id
	ManagementAccountId string `json:"management_account_id"`

	// 管理员或者委托管理员项目id
	ManagementProjectId string `json:"management_project_id"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 日志汇聚配置
	LogMappingConfig *[]LogMappingConfig `json:"log_mapping_config,omitempty"`
}

func (o UpdatelogConvergeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatelogConvergeConfig struct{}"
	}

	return strings.Join([]string{"UpdatelogConvergeConfig", string(data)}, " ")
}
