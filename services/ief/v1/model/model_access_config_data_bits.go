package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RTU传输模式下数据位配置，value值字段包括5、6、7、8
type AccessConfigDataBits struct {
	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~

	Value string `json:"value"`
	// 标识属性是否可选，默认为true

	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o AccessConfigDataBits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigDataBits struct{}"
	}

	return strings.Join([]string{"AccessConfigDataBits", string(data)}, " ")
}
