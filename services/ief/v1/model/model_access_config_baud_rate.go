package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RTU传输模式下波特率配置，value值字段可选50、75、110、150、200、300、1200、1800、2400、9600、19200、38400、57600、115200
type AccessConfigBaudRate struct {
	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~

	Value string `json:"value"`
	// 标识属性是否可选，默认为true

	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o AccessConfigBaudRate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigBaudRate struct{}"
	}

	return strings.Join([]string{"AccessConfigBaudRate", string(data)}, " ")
}
