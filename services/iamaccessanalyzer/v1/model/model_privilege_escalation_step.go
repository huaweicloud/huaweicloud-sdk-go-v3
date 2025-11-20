package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivilegeEscalationStep 提权访问分析路径的某个步骤。
type PrivilegeEscalationStep struct {
	Principal *FindingPrincipal `json:"principal"`

	// 本步骤涉及到的资源。
	Resources []string `json:"resources"`

	// 本步骤涉及到的操作。
	Action string `json:"action"`
}

func (o PrivilegeEscalationStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivilegeEscalationStep struct{}"
	}

	return strings.Join([]string{"PrivilegeEscalationStep", string(data)}, " ")
}
