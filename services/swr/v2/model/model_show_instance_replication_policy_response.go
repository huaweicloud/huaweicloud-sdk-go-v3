package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceReplicationPolicyResponse Response Object
type ShowInstanceReplicationPolicyResponse struct {

	// 策略ID
	Id *int32 `json:"id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 策略描述
	Description *string `json:"description,omitempty"`

	SrcRegistry *ReplicationRegistry `json:"src_registry,omitempty"`

	DestRegistry *ReplicationRegistry `json:"dest_registry,omitempty"`

	// 目标命名空间名，默认为空值
	DestNamespace *string `json:"dest_namespace,omitempty"`

	// 在镜像复制过程中，针对多路径镜像名称，可通过设置替换层级来优化仓库目录结构。支持的值为-1,0,1,2,3。dest_namespace_replace_count不传值，则默认为-1。假设原始镜像路径为 a/b/c/d/img，其中a为原始镜像的命名空间名称，目标命名空间为 ns，不同替换级别的效果如下： -1：替换所有级：'a/b/c/d/img' -> 'ns/img' 0：无替换：'a/b/c/d/img' -> 'ns/a/b/c/d/img' 1：替换1级：'a/b/c/d/img' -> 'ns/b/c/d/img' 2：替换2级：'a/b/c/d/img' -> 'ns/c/d/img' 3：替换3级：'a/b/c/d/img' -> 'ns/d/img'
	DestNamespaceReplaceCount *int32 `json:"dest_namespace_replace_count,omitempty"`

	// 源资源过滤器
	Filters *[]Filter `json:"filters,omitempty"`

	// repo过滤类型
	RepoScopeMode *string `json:"repo_scope_mode,omitempty"`

	Trigger *TriggerConfig `json:"trigger,omitempty"`

	// 是否覆盖
	Override *bool `json:"override,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceReplicationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceReplicationPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceReplicationPolicyResponse", string(data)}, " ")
}
