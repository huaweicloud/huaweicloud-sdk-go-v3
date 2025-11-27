package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsConstraintSpec struct {
	Constraint *Constraint `json:"constraint,omitempty"`

	// 约束模板ID
	ConstraintTemplateID *string `json:"constraintTemplateID,omitempty"`

	// 用户的domainID
	DomainID *string `json:"domainID,omitempty"`

	// 策略实例下发范围，当前cluster和fleet二选一
	TargetScope *string `json:"targetScope,omitempty"`

	// 策略实例下发对象，当前为clusterID或clustergroupID
	TargetID *string `json:"targetID,omitempty"`
}

func (o UcsConstraintSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsConstraintSpec struct{}"
	}

	return strings.Join([]string{"UcsConstraintSpec", string(data)}, " ")
}
