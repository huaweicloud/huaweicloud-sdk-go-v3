package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StepRunInputs struct {

	// **参数解释**： 输入参数名。 **取值范围**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 输入参数值。 **取值范围**： 不涉及。
	Value *interface{} `json:"value,omitempty"`
}

func (o StepRunInputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepRunInputs struct{}"
	}

	return strings.Join([]string{"StepRunInputs", string(data)}, " ")
}
