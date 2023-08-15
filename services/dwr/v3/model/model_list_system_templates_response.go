package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemTemplatesResponse Response Object
type ListSystemTemplatesResponse struct {

	// 列表条目数。
	Count *int32 `json:"count,omitempty"`

	// 模板列表信息。
	ActionTemplates *[]ActionTemplateItem `json:"action_templates,omitempty"`

	// 下一次查询的起始位置。 下一次查询的起始位置。
	Offset *int32 `json:"offset,omitempty"`

	// 如果本次没有返回全部结果，响应请求中将包含此字段，用于标明本次请求列举到的最后一个算子。如果is_truncated为false，该字段不会返回。
	IsTruncated *bool `json:"is_truncated,omitempty"`

	// 查询到符合条件的列表总条数。
	Total *int32 `json:"total,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSystemTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListSystemTemplatesResponse", string(data)}, " ")
}
