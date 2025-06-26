package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReplicationPolicy struct {

	// 策略ID
	Id *int32 `json:"id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 策略描述描述
	Description *string `json:"description,omitempty"`

	SrcRegistry *ReplicationRegistry `json:"src_registry,omitempty"`

	DestRegistry *ReplicationRegistry `json:"dest_registry,omitempty"`

	// 目标命名空间名，默认为空值
	DestNamespace *string `json:"dest_namespace,omitempty"`

	// 源资源过滤器
	Filters *[]Filter `json:"filters,omitempty"`

	// repo的范围模式
	RepoScopeMode *string `json:"repo_scope_mode,omitempty"`

	Trigger *TriggerConfig `json:"trigger,omitempty"`

	// 是否覆盖
	Override *bool `json:"override,omitempty"`

	// 是否使用
	Enabled *bool `json:"enabled,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ReplicationPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationPolicy struct{}"
	}

	return strings.Join([]string{"ReplicationPolicy", string(data)}, " ")
}
