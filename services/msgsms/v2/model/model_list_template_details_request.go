package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplateDetailsRequest struct {

	// 应用key
	AppKey *string `json:"app_key,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 流程状态 1. Adopted: 通过 2. Reviewing: 审核中 3. Reject: 拒绝
	FlowStatus *string `json:"flow_status,omitempty"`

	// 是否存在变量
	HasVariable *string `json:"has_variable,omitempty"`

	// 地域 1. cn：国内 2. intl：国际
	Region *string `json:"region,omitempty"`

	// 签名名称
	SignName *string `json:"sign_name,omitempty"`

	// 排序方式 1. desc：降序 2. asc：升序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 模板类型 1. VERIFY_CODE_TYPE: 验证码类 2. PROMOTION_TYPE: 推广类 3. NOTIFY_TYPE: 通知类
	TemplateType *string `json:"template_type,omitempty"`
}

func (o ListTemplateDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateDetailsRequest", string(data)}, " ")
}
