package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateReplicationPolicyRequestBody struct {

	// 策略名称, 1-256字符组成，只能包含英文大小写、数字、汉字、中划线和下划线。
	Name string `json:"name"`

	// 策略描述
	Description *string `json:"description,omitempty"`

	SrcRegistry *ReplicationRegistry `json:"src_registry,omitempty"`

	DestRegistry *ReplicationRegistry `json:"dest_registry,omitempty"`

	// 目标命名空间名，默认为空值
	DestNamespace *string `json:"dest_namespace,omitempty"`

	// 过滤器
	Filters []Filter `json:"filters"`

	// repo的选择方式，regular，selection
	RepoScopeMode string `json:"repo_scope_mode"`

	Trigger *TriggerConfig `json:"trigger"`

	// 是否覆盖，默认为false
	Override *bool `json:"override,omitempty"`

	// 是否启用，默认为false
	Enabled bool `json:"enabled"`
}

func (o UpdateReplicationPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplicationPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateReplicationPolicyRequestBody", string(data)}, " ")
}
