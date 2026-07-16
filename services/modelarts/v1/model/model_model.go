package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Model 模型信息
type Model struct {

	// **参数解释**：模型名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：模型OBS路径。 **取值范围**：不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释**：量化数据类型。 **取值范围**：- w8A8 - fp16
	QuantType *string `json:"quant_type,omitempty"`
}

func (o Model) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Model struct{}"
	}

	return strings.Join([]string{"Model", string(data)}, " ")
}
