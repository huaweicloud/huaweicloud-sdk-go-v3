package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyUpdate struct {

	// 是否启用策略
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// 策略名称
	Name *string `json:"name,omitempty" xml:"name"`

	OperationDefinition *PolicyoOdCreate `json:"operation_definition,omitempty" xml:"operation_definition"`

	Trigger *PolicyTriggerReq `json:"trigger,omitempty" xml:"trigger"`
}

func (o PolicyUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyUpdate struct{}"
	}

	return strings.Join([]string{"PolicyUpdate", string(data)}, " ")
}
