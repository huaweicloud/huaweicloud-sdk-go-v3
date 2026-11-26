package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperatorArgItem **参数解释**：算子的业务参数条目，当参数类型为enum时单选，参数类型为list时多选。
type OperatorArgItem struct {

	// **参数解释**：参数名。
	Name string `json:"name"`

	// **参数解释**：参数值
	Value string `json:"value"`

	// **参数解释**：参数说明
	Tips *string `json:"tips,omitempty"`
}

func (o OperatorArgItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperatorArgItem struct{}"
	}

	return strings.Join([]string{"OperatorArgItem", string(data)}, " ")
}
