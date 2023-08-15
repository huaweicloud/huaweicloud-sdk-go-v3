package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMyActionTemplateResponse Response Object
type ListMyActionTemplateResponse struct {

	// 列表条目数
	Count *int32 `json:"count,omitempty"`

	// 模板列表信息
	ActionTemplates *[]ThirdActionTemplateItem `json:"action_templates,omitempty"`

	// 下一次查询的起始位置。
	Offset *int32 `json:"offset,omitempty"`

	// 表明是否本次返回的结果列表被截断。true：表示本次没有返回全部结果。false：表示本次已经返回了全部结果。
	IsTruncated *bool `json:"is_truncated,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMyActionTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMyActionTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListMyActionTemplateResponse", string(data)}, " ")
}
