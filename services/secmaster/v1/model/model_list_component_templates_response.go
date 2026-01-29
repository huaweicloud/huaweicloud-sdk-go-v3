package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentTemplatesResponse Response Object
type ListComponentTemplatesResponse struct {

	// **参数解释**: 响应的返回码 **约束限制**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 数据总条数 **约束限制**: 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 当前页大小 **约束限制**: 不涉及
	Size *int32 `json:"size,omitempty"`

	// **参数解释**: 当前页码 **约束限制**: 不涉及
	Page *int32 `json:"page,omitempty"`

	// **参数解释**: 响应的错误信息 **约束限制**: 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**: 是否响应成功 **约束限制**: 不涉及
	Success *bool `json:"success,omitempty"`

	// **参数解释**: 请求id **约束限制**: 不涉及
	RequestId *string `json:"request_id,omitempty"`

	// 插件配置模板列表
	Data           *[]TemplateInfo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListComponentTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListComponentTemplatesResponse", string(data)}, " ")
}
