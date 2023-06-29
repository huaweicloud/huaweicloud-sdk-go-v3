package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditionItem Condition define
type ConditionItem struct {

	// Name of the condition.
	Name *string `json:"name,omitempty"`

	// Detail of the condition.
	Detail *string `json:"detail,omitempty"`

	// Detail of the condition.
	Data *[]string `json:"data,omitempty"`
}

func (o ConditionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionItem struct{}"
	}

	return strings.Join([]string{"ConditionItem", string(data)}, " ")
}
