package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsConstraintRequest struct {

	// 策略模板的ID
	ConstraintTemplateID *string `json:"constraintTemplateID,omitempty"`

	// 策略执行方式，当前支持warn和deny
	EnforcementAction *string `json:"enforcementAction,omitempty"`

	// 生效的命名空间
	Namespaces *[]string `json:"namespaces,omitempty"`

	// 策略实例的参数
	Parameters *interface{} `json:"parameters,omitempty"`
}

func (o UcsConstraintRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsConstraintRequest struct{}"
	}

	return strings.Join([]string{"UcsConstraintRequest", string(data)}, " ")
}
