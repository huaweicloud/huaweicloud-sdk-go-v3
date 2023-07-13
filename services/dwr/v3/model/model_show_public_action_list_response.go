package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicActionListResponse Response Object
type ShowPublicActionListResponse struct {

	// 列表条目数
	Count *int32 `json:"count,omitempty"`

	// 模板列表信息
	ActionTemplates *[]PublicTemplateItem `json:"action_templates,omitempty"`

	// 列表总条目数
	Total *int32 `json:"total,omitempty"`

	// 是否为分页返回
	IsTruncated *bool `json:"is_truncated,omitempty"`

	// 下次查询的起始位置
	Offset *int32 `json:"offset,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPublicActionListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicActionListResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicActionListResponse", string(data)}, " ")
}
