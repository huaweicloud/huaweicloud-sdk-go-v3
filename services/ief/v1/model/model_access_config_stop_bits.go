package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RTU传输模式下停止位配置，value值字段包含1、2
type AccessConfigStopBits struct {
	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~

	Value string `json:"value"`
	// 标识属性是否可选，默认为true

	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o AccessConfigStopBits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigStopBits struct{}"
	}

	return strings.Join([]string{"AccessConfigStopBits", string(data)}, " ")
}
