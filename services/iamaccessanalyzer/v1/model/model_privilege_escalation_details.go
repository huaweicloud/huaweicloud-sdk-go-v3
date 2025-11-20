package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivilegeEscalationDetails 提权访问分析详细结果。
type PrivilegeEscalationDetails struct {

	// 指定的待分析操作集合。
	Actions []string `json:"actions"`

	// 资源的唯一资源标识符。
	Resource string `json:"resource"`

	Principal *FindingPrincipal `json:"principal"`

	// 能够通过提权访问路径触发的操作代表。
	ActiveAction string `json:"active_action"`

	Path []PrivilegeEscalationStep `json:"path"`
}

func (o PrivilegeEscalationDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivilegeEscalationDetails struct{}"
	}

	return strings.Join([]string{"PrivilegeEscalationDetails", string(data)}, " ")
}
