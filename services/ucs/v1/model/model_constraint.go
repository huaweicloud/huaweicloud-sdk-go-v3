package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Constraint struct {

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *ConstraintSpec `json:"spec,omitempty"`
}

func (o Constraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Constraint struct{}"
	}

	return strings.Join([]string{"Constraint", string(data)}, " ")
}
