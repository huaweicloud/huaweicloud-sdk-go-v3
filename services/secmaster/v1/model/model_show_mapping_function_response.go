package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMappingFunctionResponse Response Object
type ShowMappingFunctionResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// 分类映射函数数据
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowMappingFunctionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMappingFunctionResponse struct{}"
	}

	return strings.Join([]string{"ShowMappingFunctionResponse", string(data)}, " ")
}
