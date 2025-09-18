package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputRespOutputResult struct {

	// **参数解释**： 步骤输出key值。 **取值范围**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 步骤输出value值。 **取值范围**： 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o OutputRespOutputResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputRespOutputResult struct{}"
	}

	return strings.Join([]string{"OutputRespOutputResult", string(data)}, " ")
}
