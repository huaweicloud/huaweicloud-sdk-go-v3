package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatusViolation struct {

	// 违规资源类型
	Kind *string `json:"kind,omitempty"`

	// 违规资源名称
	Name *string `json:"name,omitempty"`

	// 违规资源所在命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 详细违规信息
	Message *string `json:"message,omitempty"`

	// 审计行为
	EnforcementAction *string `json:"enforcementAction,omitempty"`
}

func (o StatusViolation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusViolation struct{}"
	}

	return strings.Join([]string{"StatusViolation", string(data)}, " ")
}
