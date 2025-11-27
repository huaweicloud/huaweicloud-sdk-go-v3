package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsConstraint struct {

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *UcsConstraintSpec `json:"spec,omitempty"`

	Status *UcsConstraintStatus `json:"status,omitempty"`
}

func (o UcsConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsConstraint struct{}"
	}

	return strings.Join([]string{"UcsConstraint", string(data)}, " ")
}
