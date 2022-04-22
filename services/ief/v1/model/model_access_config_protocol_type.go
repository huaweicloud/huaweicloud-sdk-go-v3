package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设备访问类型，包含： - userdefine：自定义访问协议； - TCP：modbus访问协议中的一种访问形式； - RTU：modbus访问协议中的一种访问形式； - opc-ua：opc-ua访问协议；
type AccessConfigProtocolType struct {

	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~
	Value string `json:"value"`

	// 标识属性是否可选，默认为true
	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o AccessConfigProtocolType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigProtocolType struct{}"
	}

	return strings.Join([]string{"AccessConfigProtocolType", string(data)}, " ")
}
