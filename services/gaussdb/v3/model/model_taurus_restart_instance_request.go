package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaurusRestartInstanceRequest 重启实例
type TaurusRestartInstanceRequest struct {

	// 实例是否延迟重启，默认false，立即重启。  - true: 延迟重启，实例将在运维时间窗内自动重启。 - false: 立即重启。
	Delay *bool `json:"delay,omitempty"`

	// **参数解释**：  实例重启后生效的max_connections内核参数配置值。max_connections用于设置允许同时连接到数据库实例的最大客户端连接数。  **约束限制**：  不涉及。  **取值范围**： 30-100000。  **默认取值**： 不涉及。
	MaxConnections *int32 `json:"max_connections,omitempty"`
}

func (o TaurusRestartInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusRestartInstanceRequest struct{}"
	}

	return strings.Join([]string{"TaurusRestartInstanceRequest", string(data)}, " ")
}
