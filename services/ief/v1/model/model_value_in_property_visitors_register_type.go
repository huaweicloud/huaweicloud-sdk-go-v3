package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 寄存器类型，value值字段可选为CoilsRegisters、HoldingRegisters、DiscreteInputsRegisters、InputRegisters
type ValueInPropertyVisitorsRegisterType struct {
	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~

	Value string `json:"value"`
	// 标识属性是否可选，默认为true

	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o ValueInPropertyVisitorsRegisterType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInPropertyVisitorsRegisterType struct{}"
	}

	return strings.Join([]string{"ValueInPropertyVisitorsRegisterType", string(data)}, " ")
}
