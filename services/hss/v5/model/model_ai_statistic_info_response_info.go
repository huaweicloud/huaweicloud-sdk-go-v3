package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiStatisticInfoResponseInfo AI组件的统计列表信息
type AiStatisticInfoResponseInfo struct {

	// **参数解释**： AI组件对应类型的名称 **取值范围**： 字符长度1-256位
	AiComponentName *string `json:"ai_component_name,omitempty"`

	// **参数解释**： AI组件所在的服务器数量
	Num *int32 `json:"num,omitempty"`
}

func (o AiStatisticInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiStatisticInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"AiStatisticInfoResponseInfo", string(data)}, " ")
}
