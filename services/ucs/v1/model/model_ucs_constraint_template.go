package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsConstraintTemplate struct {

	// API类型，固定值“ConstraintTemplate”
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	// 约束模板的详细属性
	Spec *interface{} `json:"spec,omitempty"`
}

func (o UcsConstraintTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsConstraintTemplate struct{}"
	}

	return strings.Join([]string{"UcsConstraintTemplate", string(data)}, " ")
}
